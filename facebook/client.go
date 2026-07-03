package facebook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
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

type Option func(*Client)

type Client struct {
	AccessToken string
	APIVersion  string
	BaseURL     string
	HTTPClient  *http.Client
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

type Cursor[T any] struct {
	Data    []T                    `json:"data"`
	Paging  *Paging                `json:"paging,omitempty"`
	Summary map[string]interface{} `json:"summary,omitempty"`
}

type Paging struct {
	Cursors  map[string]string `json:"cursors,omitempty"`
	Next     string            `json:"next,omitempty"`
	Previous string            `json:"previous,omitempty"`
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
}

func (err *APIError) Error() string {
	if err.ErrorInfo.Message != "" {
		return fmt.Sprintf("facebook API error: status=%d code=%d message=%s", err.StatusCode, err.ErrorInfo.Code, err.ErrorInfo.Message)
	}
	return fmt.Sprintf("facebook API error: status=%d body=%s", err.StatusCode, err.RawBody)
}

func (client *Client) Request(ctx context.Context, method string, graphPath string, params Params, out interface{}) error {
	req, err := client.newRequest(ctx, method, graphPath, params)
	if err != nil {
		return err
	}

	httpClient := client.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		var payload struct {
			Error *ErrorResponse `json:"error"`
		}
		if json.Unmarshal(body, &payload) == nil && payload.Error != nil {
			return &APIError{
				StatusCode: resp.StatusCode,
				ErrorInfo:  *payload.Error,
				RawBody:    string(body),
			}
		}
		return &APIError{StatusCode: resp.StatusCode, RawBody: string(body)}
	}

	if out == nil || len(bytes.TrimSpace(body)) == 0 {
		return nil
	}

	return json.Unmarshal(body, out)
}

func (client *Client) newRequest(ctx context.Context, method string, graphPath string, params Params) (*http.Request, error) {
	endpoint, err := url.Parse(client.endpointURL(graphPath))
	if err != nil {
		return nil, err
	}

	values, files, err := encodeParams(params)
	if err != nil {
		return nil, err
	}
	if client.AccessToken != "" && values.Get("access_token") == "" {
		values.Set("access_token", client.AccessToken)
	}

	if method == http.MethodGet || method == http.MethodDelete {
		endpoint.RawQuery = values.Encode()
		return http.NewRequestWithContext(ctx, method, endpoint.String(), nil)
	}

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
		if value == nil {
			continue
		}

		switch typed := value.(type) {
		case FileParam:
			files[key] = typed
			continue
		case *FileParam:
			if typed != nil {
				files[key] = *typed
			}
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
	switch typed := value.(type) {
	case string:
		return typed, nil
	case ID:
		return string(typed), nil
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
