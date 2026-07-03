package facebook

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultBaseURL    = "https://graph.facebook.com"
	DefaultAPIVersion = "v25.0"
)

type Params map[string]interface{}

func Ptr[T any](value T) *T {
	return &value
}

func Fields(fields ...string) Params {
	return Params{"fields": fields}
}

func (params Params) With(key string, value interface{}) Params {
	out := Params{}
	for currentKey, currentValue := range params {
		out[currentKey] = currentValue
	}
	out[key] = value
	return out
}

func (params Params) WithFields(fields ...string) Params {
	return params.With("fields", fields)
}

type Option func(*Client)

type RetryPolicy struct {
	MaxAttempts        int
	InitialDelay       time.Duration
	MaxDelay           time.Duration
	Statuses           map[int]bool
	RetryUnsafeMethods bool
}

type Client struct {
	AccessToken string
	AppSecret   string
	APIVersion  string
	BaseURL     string
	HTTPClient  *http.Client
	RetryPolicy *RetryPolicy
}

func NewClient(accessToken string, options ...Option) *Client {
	client := &Client{
		AccessToken: accessToken,
		APIVersion:  DefaultAPIVersion,
		BaseURL:     DefaultBaseURL,
		HTTPClient:  http.DefaultClient,
	}

	for _, option := range options {
		option(client)
	}

	return client
}

func WithAPIVersion(apiVersion string) Option {
	return func(client *Client) {
		client.APIVersion = apiVersion
	}
}

func WithBaseURL(baseURL string) Option {
	return func(client *Client) {
		client.BaseURL = baseURL
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(client *Client) {
		client.HTTPClient = httpClient
	}
}

func WithAppSecret(appSecret string) Option {
	return func(client *Client) {
		client.AppSecret = appSecret
	}
}

func WithRetry(maxAttempts int) Option {
	policy := DefaultRetryPolicy()
	policy.MaxAttempts = maxAttempts
	return WithRetryPolicy(policy)
}

func WithRetryPolicy(policy RetryPolicy) Option {
	return func(client *Client) {
		client.RetryPolicy = &policy
	}
}

func WithRetryNonIdempotentRequests() Option {
	return func(client *Client) {
		if client.RetryPolicy == nil {
			policy := DefaultRetryPolicy()
			client.RetryPolicy = &policy
		}
		client.RetryPolicy.RetryUnsafeMethods = true
	}
}

func DefaultRetryPolicy() RetryPolicy {
	return RetryPolicy{
		MaxAttempts:  3,
		InitialDelay: 500 * time.Millisecond,
		MaxDelay:     8 * time.Second,
		Statuses: map[int]bool{
			http.StatusTooManyRequests:     true,
			http.StatusInternalServerError: true,
			http.StatusBadGateway:          true,
			http.StatusServiceUnavailable:  true,
			http.StatusGatewayTimeout:      true,
		},
	}
}

type ID string

func (id ID) String() string {
	return string(id)
}

func (id *ID) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		*id = ""
		return nil
	}

	var value string
	if len(data) > 0 && data[0] == '"' {
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		*id = ID(value)
		return nil
	}

	*id = ID(string(data))
	return nil
}

type FileParam struct {
	FileName string
	Reader   io.Reader
}

type Time struct {
	time.Time
}

func (value *Time) UnmarshalJSON(data []byte) error {
	parsed, err := parseFacebookTime(data)
	if err != nil {
		return err
	}
	value.Time = parsed
	return nil
}

func (value Time) MarshalJSON() ([]byte, error) {
	if value.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(value.UTC().Format(time.RFC3339))
}

func (value Time) String() string {
	if value.IsZero() {
		return ""
	}
	return value.UTC().Format(time.RFC3339)
}

type Date struct {
	time.Time
}

func (value *Date) UnmarshalJSON(data []byte) error {
	parsed, err := parseFacebookDate(data)
	if err != nil {
		return err
	}
	value.Time = parsed
	return nil
}

func (value Date) MarshalJSON() ([]byte, error) {
	if value.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(value.Format("2006-01-02"))
}

func (value Date) String() string {
	if value.IsZero() {
		return ""
	}
	return value.Format("2006-01-02")
}

type Response struct {
	StatusCode           int
	Header               http.Header
	Body                 []byte
	TraceID              string
	AppUsage             map[string]interface{}
	AdAccountUsage       map[string]interface{}
	BusinessUseCaseUsage map[string]interface{}
}

func (response *Response) HeaderValue(name string) string {
	if response == nil {
		return ""
	}
	return response.Header.Get(name)
}

func (response *Response) BodyString() string {
	if response == nil {
		return ""
	}
	return string(response.Body)
}

type Cursor[T any] struct {
	Data    []T                    `json:"data"`
	Paging  *Paging                `json:"paging,omitempty"`
	Summary map[string]interface{} `json:"summary,omitempty"`
	client  *Client
}

type Paging struct {
	Cursors  map[string]string `json:"cursors,omitempty"`
	Next     string            `json:"next,omitempty"`
	Previous string            `json:"previous,omitempty"`
}

func (cursor *Cursor[T]) HasNext() bool {
	return cursor != nil && cursor.Paging != nil && cursor.Paging.Next != ""
}

func (cursor *Cursor[T]) HasPrevious() bool {
	return cursor != nil && cursor.Paging != nil && cursor.Paging.Previous != ""
}

func (cursor *Cursor[T]) Next(ctx context.Context) (*Cursor[T], error) {
	if !cursor.HasNext() {
		return nil, nil
	}
	return cursor.page(ctx, cursor.Paging.Next)
}

func (cursor *Cursor[T]) Previous(ctx context.Context) (*Cursor[T], error) {
	if !cursor.HasPrevious() {
		return nil, nil
	}
	return cursor.page(ctx, cursor.Paging.Previous)
}

func (cursor *Cursor[T]) All(ctx context.Context) ([]T, error) {
	if cursor == nil {
		return nil, nil
	}
	items := []T{}
	err := cursor.Each(ctx, func(item T) error {
		items = append(items, item)
		return nil
	})
	return items, err
}

func (cursor *Cursor[T]) Each(ctx context.Context, fn func(T) error) error {
	if cursor == nil || fn == nil {
		return nil
	}
	return cursor.Pages(ctx, func(page *Cursor[T]) error {
		for _, item := range page.Data {
			if err := fn(item); err != nil {
				return err
			}
		}
		return nil
	})
}

func (cursor *Cursor[T]) Pages(ctx context.Context, fn func(*Cursor[T]) error) error {
	if cursor == nil || fn == nil {
		return nil
	}
	for current := cursor; current != nil; {
		if err := fn(current); err != nil {
			return err
		}
		if !current.HasNext() {
			return nil
		}
		next, err := current.Next(ctx)
		if err != nil {
			return err
		}
		current = next
	}
	return nil
}

func (cursor *Cursor[T]) page(ctx context.Context, pageURL string) (*Cursor[T], error) {
	if cursor == nil || cursor.client == nil {
		return nil, fmt.Errorf("facebook cursor has nil client")
	}
	var out Cursor[T]
	err := cursor.client.RequestURL(ctx, http.MethodGet, pageURL, nil, &out)
	return &out, err
}

func (cursor *Cursor[T]) setClient(client *Client) {
	if cursor != nil {
		cursor.client = client
	}
}

type ErrorResponse struct {
	Code           int    `json:"code,omitempty"`
	ErrorSubcode   int    `json:"error_subcode,omitempty"`
	FBTraceID      string `json:"fbtrace_id,omitempty"`
	Message        string `json:"message,omitempty"`
	Type           string `json:"type,omitempty"`
	IsTransient    bool   `json:"is_transient,omitempty"`
	ErrorUserMsg   string `json:"error_user_msg,omitempty"`
	ErrorUserTitle string `json:"error_user_title,omitempty"`
}

type APIError struct {
	StatusCode int
	ErrorInfo  ErrorResponse
	RawBody    string
	Response   *Response
}

func (err *APIError) Error() string {
	if err.ErrorInfo.Message != "" {
		return fmt.Sprintf("facebook API error: status=%d code=%d message=%s", err.StatusCode, err.ErrorInfo.Code, err.ErrorInfo.Message)
	}
	return fmt.Sprintf("facebook API error: status=%d body=%s", err.StatusCode, err.RawBody)
}

func AsAPIError(err error) (*APIError, bool) {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
}

func IsRateLimit(err error) bool {
	apiErr, ok := AsAPIError(err)
	if !ok {
		return false
	}
	if apiErr.StatusCode == http.StatusTooManyRequests {
		return true
	}
	switch apiErr.ErrorInfo.Code {
	case 4, 17, 32, 613:
		return true
	}
	return false
}

func IsOAuthError(err error) bool {
	apiErr, ok := AsAPIError(err)
	if !ok {
		return false
	}
	return apiErr.StatusCode == http.StatusUnauthorized || apiErr.ErrorInfo.Code == 190 || apiErr.ErrorInfo.Type == "OAuthException"
}

func IsPermissionError(err error) bool {
	apiErr, ok := AsAPIError(err)
	if !ok {
		return false
	}
	if apiErr.StatusCode == http.StatusForbidden {
		return true
	}
	switch apiErr.ErrorInfo.Code {
	case 10, 200, 294:
		return true
	}
	return false
}

func IsTransient(err error) bool {
	apiErr, ok := AsAPIError(err)
	if !ok {
		return false
	}
	return apiErr.ErrorInfo.IsTransient || apiErr.StatusCode == http.StatusTooManyRequests || apiErr.StatusCode >= http.StatusInternalServerError
}

func StatusCode(err error) int {
	apiErr, ok := AsAPIError(err)
	if !ok {
		return 0
	}
	return apiErr.StatusCode
}

func TraceID(err error) string {
	apiErr, ok := AsAPIError(err)
	if !ok {
		return ""
	}
	if apiErr.ErrorInfo.FBTraceID != "" {
		return apiErr.ErrorInfo.FBTraceID
	}
	if apiErr.Response != nil {
		return apiErr.Response.TraceID
	}
	return ""
}

func (client *Client) Request(ctx context.Context, method string, graphPath string, params Params, out interface{}) error {
	_, err := client.RequestWithResponse(ctx, method, graphPath, params, out)
	return err
}

func (client *Client) RequestWithResponse(ctx context.Context, method string, graphPath string, params Params, out interface{}) (*Response, error) {
	return client.doRequest(ctx, method, params, out, func() (*http.Request, error) {
		return client.newRequest(ctx, method, graphPath, params)
	})
}

func (client *Client) RequestURL(ctx context.Context, method string, requestURL string, params Params, out interface{}) error {
	_, err := client.RequestURLWithResponse(ctx, method, requestURL, params, out)
	return err
}

func (client *Client) RequestURLWithResponse(ctx context.Context, method string, requestURL string, params Params, out interface{}) (*Response, error) {
	return client.doRequest(ctx, method, params, out, func() (*http.Request, error) {
		return client.newRequestURL(ctx, method, requestURL, params)
	})
}

func (client *Client) doRequest(ctx context.Context, method string, params Params, out interface{}, newRequest func() (*http.Request, error)) (*Response, error) {
	if client == nil {
		return nil, fmt.Errorf("facebook client is nil")
	}

	policy := client.effectiveRetryPolicy()
	if paramsHaveFiles(params) || !canRetryMethod(method, policy) {
		policy.MaxAttempts = 1
	}

	var response *Response
	var lastErr error
	for attempt := 1; attempt <= policy.MaxAttempts; attempt++ {
		req, err := newRequest()
		if err != nil {
			return nil, err
		}

		resp, body, err := client.doHTTP(req)
		response = newResponse(resp, body)
		if err != nil {
			lastErr = err
		} else {
			lastErr = nil
		}

		if !shouldRetry(policy, resp, body, err) || attempt == policy.MaxAttempts {
			if err != nil {
				return response, err
			}
			return client.decodeResponse(resp, body, out)
		}

		if err := sleepWithContext(ctx, retryDelay(policy, attempt, resp)); err != nil {
			return response, err
		}
	}
	return response, lastErr
}

func (client *Client) doHTTP(req *http.Request) (*http.Response, []byte, error) {
	httpClient := client.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, nil, err
	}
	return resp, body, nil
}

func newResponse(resp *http.Response, body []byte) *Response {
	if resp == nil {
		return nil
	}
	header := resp.Header.Clone()
	return &Response{
		StatusCode:           resp.StatusCode,
		Header:               header,
		Body:                 append([]byte(nil), body...),
		TraceID:              header.Get("x-fb-trace-id"),
		AppUsage:             decodeHeaderObject(header, "x-app-usage"),
		AdAccountUsage:       decodeHeaderObject(header, "x-ad-account-usage"),
		BusinessUseCaseUsage: decodeHeaderObject(header, "x-business-use-case-usage"),
	}
}

func decodeHeaderObject(header http.Header, name string) map[string]interface{} {
	value := strings.TrimSpace(header.Get(name))
	if value == "" {
		return nil
	}
	out := map[string]interface{}{}
	if json.Unmarshal([]byte(value), &out) != nil {
		return nil
	}
	return out
}

func (client *Client) decodeResponse(resp *http.Response, body []byte, out interface{}) (*Response, error) {
	response := newResponse(resp, body)
	if resp == nil {
		return response, fmt.Errorf("facebook response is nil")
	}
	if resp.StatusCode >= http.StatusBadRequest {
		var payload struct {
			Error *ErrorResponse `json:"error"`
		}
		if json.Unmarshal(body, &payload) == nil && payload.Error != nil {
			return response, &APIError{
				StatusCode: resp.StatusCode,
				ErrorInfo:  *payload.Error,
				RawBody:    string(body),
				Response:   response,
			}
		}
		return response, &APIError{StatusCode: resp.StatusCode, RawBody: string(body), Response: response}
	}

	if out == nil || len(bytes.TrimSpace(body)) == 0 {
		return response, nil
	}
	if err := json.Unmarshal(body, out); err != nil {
		return response, err
	}
	client.attachClient(out)
	return response, nil
}

func (client *Client) attachClient(out interface{}) {
	if setter, ok := out.(interface{ setClient(*Client) }); ok {
		setter.setClient(client)
	}
}

func (client *Client) newRequest(ctx context.Context, method string, graphPath string, params Params) (*http.Request, error) {
	endpoint, err := url.Parse(client.endpointURL(graphPath))
	if err != nil {
		return nil, err
	}
	return client.newRequestForEndpoint(ctx, method, endpoint, params)
}

func (client *Client) newRequestURL(ctx context.Context, method string, requestURL string, params Params) (*http.Request, error) {
	endpoint, err := url.Parse(requestURL)
	if err != nil {
		return nil, err
	}
	if endpoint.Scheme == "" || endpoint.Host == "" {
		endpoint, err = url.Parse(client.endpointURL(requestURL))
		if err != nil {
			return nil, err
		}
	}
	return client.newRequestForEndpoint(ctx, method, endpoint, params)
}

func (client *Client) newRequestForEndpoint(ctx context.Context, method string, endpoint *url.URL, params Params) (*http.Request, error) {
	values, files, err := encodeParams(params)
	if err != nil {
		return nil, err
	}

	if method == http.MethodGet || method == http.MethodDelete {
		query := endpoint.Query()
		for key, items := range values {
			for _, value := range items {
				query.Add(key, value)
			}
		}
		client.addAuthParams(query)
		endpoint.RawQuery = query.Encode()
		return http.NewRequestWithContext(ctx, method, endpoint.String(), nil)
	}

	client.addAuthParams(values)
	if len(files) > 0 {
		body, contentType, err := buildMultipartBody(values, files)
		if err != nil {
			return nil, err
		}
		req, err := http.NewRequestWithContext(ctx, method, endpoint.String(), body)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", contentType)
		return req, nil
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint.String(), strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

func (client *Client) addAuthParams(values url.Values) {
	accessToken := values.Get("access_token")
	if accessToken == "" && client.AccessToken != "" {
		accessToken = client.AccessToken
		values.Set("access_token", accessToken)
	}
	if accessToken != "" && client.AppSecret != "" && values.Get("appsecret_proof") == "" {
		values.Set("appsecret_proof", appSecretProof(accessToken, client.AppSecret))
	}
}

func appSecretProof(accessToken string, appSecret string) string {
	mac := hmac.New(sha256.New, []byte(appSecret))
	mac.Write([]byte(accessToken))
	return hex.EncodeToString(mac.Sum(nil))
}

func parseFacebookTime(data []byte) (time.Time, error) {
	data = bytes.TrimSpace(data)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		return time.Time{}, nil
	}
	var text string
	if len(data) > 0 && data[0] == '"' {
		if err := json.Unmarshal(data, &text); err != nil {
			return time.Time{}, err
		}
		text = strings.TrimSpace(text)
		if text == "" {
			return time.Time{}, nil
		}
		for _, layout := range []string{
			time.RFC3339Nano,
			time.RFC3339,
			"2006-01-02T15:04:05.000-0700",
			"2006-01-02T15:04:05-0700",
			"2006-01-02 15:04:05 -0700",
			"2006-01-02 15:04:05",
			"2006-01-02",
		} {
			if parsed, err := time.Parse(layout, text); err == nil {
				return parsed, nil
			}
		}
		return time.Time{}, fmt.Errorf("parse facebook time %q", text)
	}
	seconds, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return time.Time{}, err
	}
	sec := int64(seconds)
	nsec := int64((seconds - float64(sec)) * 1e9)
	return time.Unix(sec, nsec).UTC(), nil
}

func parseFacebookDate(data []byte) (time.Time, error) {
	data = bytes.TrimSpace(data)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		return time.Time{}, nil
	}
	var text string
	if len(data) > 0 && data[0] == '"' {
		if err := json.Unmarshal(data, &text); err != nil {
			return time.Time{}, err
		}
		text = strings.TrimSpace(text)
		if text == "" {
			return time.Time{}, nil
		}
		if parsed, err := time.Parse("2006-01-02", text); err == nil {
			return parsed, nil
		}
	}
	return parseFacebookTime(data)
}

func (client *Client) effectiveRetryPolicy() RetryPolicy {
	if client != nil && client.RetryPolicy != nil {
		policy := *client.RetryPolicy
		return normalizeRetryPolicy(policy)
	}
	return RetryPolicy{MaxAttempts: 1}
}

func normalizeRetryPolicy(policy RetryPolicy) RetryPolicy {
	if policy.MaxAttempts < 1 {
		policy.MaxAttempts = 1
	}
	if policy.InitialDelay < 0 {
		policy.InitialDelay = 0
	}
	if policy.MaxDelay <= 0 {
		policy.MaxDelay = policy.InitialDelay
	}
	if policy.InitialDelay > policy.MaxDelay {
		policy.InitialDelay = policy.MaxDelay
	}
	if policy.Statuses == nil {
		policy.Statuses = DefaultRetryPolicy().Statuses
	}
	return policy
}

func canRetryMethod(method string, policy RetryPolicy) bool {
	if policy.RetryUnsafeMethods {
		return true
	}
	switch strings.ToUpper(strings.TrimSpace(method)) {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		return true
	}
	return false
}

func shouldRetry(policy RetryPolicy, resp *http.Response, body []byte, err error) bool {
	if policy.MaxAttempts <= 1 {
		return false
	}
	if err != nil {
		return true
	}
	if resp == nil {
		return false
	}
	if policy.Statuses[resp.StatusCode] {
		return true
	}
	return responseIsTransient(body)
}

func responseIsTransient(body []byte) bool {
	var payload struct {
		Error *ErrorResponse `json:"error"`
	}
	return json.Unmarshal(body, &payload) == nil && payload.Error != nil && payload.Error.IsTransient
}

func retryDelay(policy RetryPolicy, attempt int, resp *http.Response) time.Duration {
	if delay, ok := retryAfterDelay(resp); ok {
		return delay
	}
	delay := policy.InitialDelay
	for i := 1; i < attempt; i++ {
		delay *= 2
		if delay >= policy.MaxDelay {
			return policy.MaxDelay
		}
	}
	return delay
}

func retryAfterDelay(resp *http.Response) (time.Duration, bool) {
	if resp == nil {
		return 0, false
	}
	value := strings.TrimSpace(resp.Header.Get("Retry-After"))
	if value == "" {
		return 0, false
	}
	if seconds, err := strconv.Atoi(value); err == nil {
		return time.Duration(seconds) * time.Second, true
	}
	if retryAt, err := http.ParseTime(value); err == nil {
		delay := time.Until(retryAt)
		if delay < 0 {
			delay = 0
		}
		return delay, true
	}
	return 0, false
}

func sleepWithContext(ctx context.Context, delay time.Duration) error {
	if delay <= 0 {
		return ctx.Err()
	}
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func paramsHaveFiles(params Params) bool {
	for _, value := range params {
		value = dereferenceParamValue(value)
		switch value.(type) {
		case FileParam:
			return true
		}
	}
	return false
}

func (client *Client) endpointURL(graphPath string) string {
	baseURL := strings.TrimRight(client.BaseURL, "/")
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	parts := []string{baseURL}
	if version := strings.Trim(client.APIVersion, "/"); version != "" {
		parts = append(parts, version)
	}
	if graphPath = strings.Trim(graphPath, "/"); graphPath != "" {
		parts = append(parts, graphPath)
	}
	return strings.Join(parts, "/")
}

func encodeParams(params Params) (url.Values, map[string]FileParam, error) {
	values := url.Values{}
	files := map[string]FileParam{}

	for key, value := range params {
		value = dereferenceParamValue(value)
		if value == nil {
			continue
		}

		switch typed := value.(type) {
		case FileParam:
			files[key] = typed
			continue
		}

		encoded, err := encodeParamValue(key, value)
		if err != nil {
			return nil, nil, fmt.Errorf("encode param %q: %w", key, err)
		}
		values.Add(key, encoded)
	}

	return values, files, nil
}

func encodeParamValue(key string, value interface{}) (string, error) {
	value = dereferenceParamValue(value)
	if value == nil {
		return "", nil
	}
	switch typed := value.(type) {
	case string:
		return typed, nil
	case ID:
		return string(typed), nil
	case Time:
		return typed.String(), nil
	case Date:
		return typed.String(), nil
	case time.Time:
		return typed.Format(time.RFC3339), nil
	case fmt.Stringer:
		return typed.String(), nil
	case []string:
		if key == "fields" {
			return strings.Join(typed, ","), nil
		}
	}

	switch value.(type) {
	case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return fmt.Sprint(value), nil
	}

	data, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func dereferenceParamValue(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	current := reflect.ValueOf(value)
	for current.IsValid() && current.Kind() == reflect.Ptr {
		if current.IsNil() {
			return nil
		}
		current = current.Elem()
	}
	if !current.IsValid() {
		return nil
	}
	return current.Interface()
}

func buildMultipartBody(values url.Values, files map[string]FileParam) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, items := range values {
		for _, value := range items {
			if err := writer.WriteField(key, value); err != nil {
				return nil, "", err
			}
		}
	}

	for key, file := range files {
		if file.Reader == nil {
			return nil, "", fmt.Errorf("file param %q has nil reader", key)
		}
		fileName := file.FileName
		if fileName == "" {
			fileName = key
		}
		part, err := writer.CreateFormFile(key, fileName)
		if err != nil {
			return nil, "", err
		}
		if _, err := io.Copy(part, file.Reader); err != nil {
			return nil, "", err
		}
	}

	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return body, writer.FormDataContentType(), nil
}

func GraphPath(parts ...string) string {
	escaped := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.Trim(part, "/")
		if part == "" {
			continue
		}
		escaped = append(escaped, url.PathEscape(part))
	}
	return "/" + strings.Join(escaped, "/")
}
