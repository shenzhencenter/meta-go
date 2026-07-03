package facebook

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Batch struct {
	client  *Client
	calls   []BatchCall
	options []BatchExecuteOption
	errors  []*BatchCallError
}

type BatchCall struct {
	Method                string
	RelativeURL           string
	Params                Params
	Body                  string
	Name                  string
	DependsOn             string
	OmitResponseOnSuccess *bool
	Headers               []BatchHeader
	decoder               func(*BatchResponse) error
}

type BatchHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type BatchResponse struct {
	Code    *int           `json:"code,omitempty"`
	Headers *[]BatchHeader `json:"headers,omitempty"`
	Body    *string        `json:"body,omitempty"`
}

type BatchOption func(*BatchCall)

type BatchExecuteOption func(*batchExecuteOptions)

type BatchRequestDecoder interface {
	BatchCall() BatchCall
	DecodeBatchResponse(*BatchResponse) error
}

type BatchRequest[T any] struct {
	call     BatchCall
	response *BatchResponse
	result   *T
	err      error
}

type AsyncBatchResponse struct {
	AsyncSessions *[]map[string]interface{} `json:"async_sessions,omitempty"`
}

type BatchCallError struct {
	Index int
	Name  string
	Err   error
}

type BatchError struct {
	Errors []*BatchCallError
}

type batchExecuteOptions struct {
	params Params
}

func NewBatch(client *Client, options ...BatchExecuteOption) *Batch {
	batch := &Batch{client: client}
	return batch.WithOptions(options...)
}

func (client *Client) NewBatch(options ...BatchExecuteOption) *Batch {
	return NewBatch(client, options...)
}

func (batch *Batch) Len() int {
	if batch == nil {
		return 0
	}
	return len(batch.calls)
}

func (batch *Batch) WithOptions(options ...BatchExecuteOption) *Batch {
	if batch == nil {
		return batch
	}
	batch.options = append(batch.options, options...)
	return batch
}

func (batch *Batch) Errors() []*BatchCallError {
	if batch == nil || len(batch.errors) == 0 {
		return nil
	}
	out := make([]*BatchCallError, len(batch.errors))
	copy(out, batch.errors)
	return out
}

func (batch *Batch) Err() error {
	if batch == nil || len(batch.errors) == 0 {
		return nil
	}
	return &BatchError{Errors: batch.Errors()}
}

func (batch *Batch) Add(method string, graphPath string, params Params, options ...BatchOption) *Batch {
	if batch == nil {
		return batch
	}
	batch.calls = append(batch.calls, NewBatchCall(method, graphPath, params, options...))
	return batch
}

func (batch *Batch) AddCall(call BatchCall, options ...BatchOption) *Batch {
	if batch == nil {
		return batch
	}
	for _, option := range options {
		option(&call)
	}
	batch.calls = append(batch.calls, call)
	return batch
}

func (batch *Batch) AddRequest(request BatchRequestDecoder, options ...BatchOption) *Batch {
	if batch == nil || request == nil {
		return batch
	}
	call := request.BatchCall()
	call.decoder = request.DecodeBatchResponse
	for _, option := range options {
		option(&call)
	}
	batch.calls = append(batch.calls, call)
	return batch
}

func NewBatchCall(method string, graphPath string, params Params, options ...BatchOption) BatchCall {
	call := BatchCall{
		Method:      method,
		RelativeURL: graphPath,
		Params:      params,
	}
	for _, option := range options {
		option(&call)
	}
	return call
}

func NewBatchRequest[T any](call BatchCall) *BatchRequest[T] {
	request := &BatchRequest[T]{call: call}
	request.call.decoder = request.DecodeBatchResponse
	return request
}

func (request *BatchRequest[T]) BatchCall() BatchCall {
	if request == nil {
		return BatchCall{}
	}
	return request.call
}

func (request *BatchRequest[T]) DecodeBatchResponse(response *BatchResponse) error {
	if request == nil {
		return nil
	}
	request.response = response
	request.result = nil
	request.err = nil
	if response == nil {
		return nil
	}
	if err := response.Err(); err != nil {
		request.err = err
		return err
	}
	var out T
	if err := response.Decode(&out); err != nil {
		request.err = err
		return err
	}
	request.result = &out
	return nil
}

func (request *BatchRequest[T]) Result() (*T, error) {
	if request == nil {
		return nil, nil
	}
	return request.result, request.err
}

func (request *BatchRequest[T]) Response() *BatchResponse {
	if request == nil {
		return nil
	}
	return request.response
}

func (request *BatchRequest[T]) Err() error {
	if request == nil {
		return nil
	}
	return request.err
}

func (batch *Batch) Execute(ctx context.Context, options ...BatchExecuteOption) ([]*BatchResponse, error) {
	if batch == nil || batch.client == nil {
		return nil, fmt.Errorf("facebook batch has nil client")
	}
	responses, callErrors, err := batch.client.executeBatch(ctx, batch.calls, batch.mergedOptions(options...)...)
	batch.errors = callErrors
	if err != nil {
		return responses, err
	}
	return responses, nil
}

func (batch *Batch) ExecuteStrict(ctx context.Context, options ...BatchExecuteOption) ([]*BatchResponse, error) {
	if batch == nil || batch.client == nil {
		return nil, fmt.Errorf("facebook batch has nil client")
	}
	responses, callErrors, err := batch.client.executeBatch(ctx, batch.calls, batch.mergedOptions(options...)...)
	batch.errors = callErrors
	if err != nil {
		return responses, err
	}
	if len(callErrors) > 0 {
		return responses, &BatchError{Errors: callErrors}
	}
	return responses, nil
}

func (batch *Batch) ExecuteAsync(ctx context.Context, options ...BatchExecuteOption) (*AsyncBatchResponse, error) {
	if batch == nil || batch.client == nil {
		return nil, fmt.Errorf("facebook batch has nil client")
	}
	return batch.client.ExecuteAsyncBatch(ctx, batch.calls, batch.mergedOptions(options...)...)
}

func (batch *Batch) mergedOptions(options ...BatchExecuteOption) []BatchExecuteOption {
	if batch == nil {
		return options
	}
	merged := make([]BatchExecuteOption, 0, len(batch.options)+len(options))
	merged = append(merged, batch.options...)
	merged = append(merged, options...)
	return merged
}

func (client *Client) ExecuteBatch(ctx context.Context, calls []BatchCall, options ...BatchExecuteOption) ([]*BatchResponse, error) {
	responses, _, err := client.executeBatch(ctx, calls, options...)
	return responses, err
}

func (client *Client) ExecuteBatchStrict(ctx context.Context, calls []BatchCall, options ...BatchExecuteOption) ([]*BatchResponse, error) {
	responses, callErrors, err := client.executeBatch(ctx, calls, options...)
	if err != nil {
		return responses, err
	}
	if len(callErrors) > 0 {
		return responses, &BatchError{Errors: callErrors}
	}
	return responses, nil
}

func (client *Client) ExecuteAsyncBatch(ctx context.Context, calls []BatchCall, options ...BatchExecuteOption) (*AsyncBatchResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("facebook client is nil")
	}
	if len(calls) == 0 {
		return nil, nil
	}

	payload, err := batchPayloadForCalls(calls)
	if err != nil {
		return nil, err
	}

	var out AsyncBatchResponse
	err = client.Request(ctx, http.MethodPost, "", batchRequestParams("asyncbatch", payload, options...), &out)
	return &out, err
}

func (client *Client) executeBatch(ctx context.Context, calls []BatchCall, options ...BatchExecuteOption) ([]*BatchResponse, []*BatchCallError, error) {
	if client == nil {
		return nil, nil, fmt.Errorf("facebook client is nil")
	}
	if len(calls) == 0 {
		return nil, nil, nil
	}

	payload, err := batchPayloadForCalls(calls)
	if err != nil {
		return nil, nil, err
	}

	var out []*BatchResponse
	err = client.Request(ctx, http.MethodPost, "", batchRequestParams("batch", payload, options...), &out)
	if err != nil {
		return out, nil, err
	}
	return out, decodeBatchResponses(calls, out), nil
}

func BatchName(name string) BatchOption {
	return func(call *BatchCall) {
		call.Name = name
	}
}

func BatchDependsOn(name string) BatchOption {
	return func(call *BatchCall) {
		call.DependsOn = name
	}
}

func BatchOmitResponseOnSuccess(value bool) BatchOption {
	return func(call *BatchCall) {
		call.OmitResponseOnSuccess = &value
	}
}

func BatchBody(body string) BatchOption {
	return func(call *BatchCall) {
		call.Body = body
	}
}

func BatchHeaders(headers ...BatchHeader) BatchOption {
	return func(call *BatchCall) {
		call.Headers = append(call.Headers, headers...)
	}
}

func BatchIncludeHeaders(value bool) BatchExecuteOption {
	return BatchParam("include_headers", value)
}

func BatchAppID(appID string) BatchExecuteOption {
	return BatchParam("batch_app_id", appID)
}

func BatchParam(key string, value interface{}) BatchExecuteOption {
	return func(options *batchExecuteOptions) {
		if key == "" {
			return
		}
		options.ensureParams()
		options.params[key] = value
	}
}

func BatchParams(params Params) BatchExecuteOption {
	return func(options *batchExecuteOptions) {
		options.ensureParams()
		for key, value := range params {
			options.params[key] = value
		}
	}
}

func BatchDecode(out interface{}) BatchOption {
	return BatchDecoder(func(response *BatchResponse) error {
		if response == nil || out == nil {
			return nil
		}
		if err := response.Err(); err != nil {
			return err
		}
		return response.Decode(out)
	})
}

func BatchDecoder(decoder func(*BatchResponse) error) BatchOption {
	return func(call *BatchCall) {
		call.decoder = decoder
	}
}

func (err *BatchCallError) Error() string {
	if err == nil || err.Err == nil {
		return ""
	}
	if err.Name != "" {
		return fmt.Sprintf("facebook batch call %d (%s): %v", err.Index, err.Name, err.Err)
	}
	return fmt.Sprintf("facebook batch call %d: %v", err.Index, err.Err)
}

func (err *BatchCallError) Unwrap() error {
	if err == nil {
		return nil
	}
	return err.Err
}

func (err *BatchError) Error() string {
	if err == nil || len(err.Errors) == 0 {
		return ""
	}
	if len(err.Errors) == 1 {
		return err.Errors[0].Error()
	}
	return fmt.Sprintf("facebook batch has %d failed calls: %v", len(err.Errors), err.Errors[0])
}

func (err *BatchError) Unwrap() error {
	if err == nil || len(err.Errors) == 0 {
		return nil
	}
	return err.Errors[0]
}

func (response *BatchResponse) IsSuccess() bool {
	if response == nil {
		return true
	}
	return response.Code != nil && *response.Code >= http.StatusOK && *response.Code < http.StatusMultipleChoices
}

func (response *BatchResponse) Decode(out interface{}) error {
	if response == nil || response.Body == nil || out == nil || strings.TrimSpace(*response.Body) == "" {
		return nil
	}
	return json.Unmarshal([]byte(*response.Body), out)
}

func (response *BatchResponse) Err() error {
	if response == nil || response.Code == nil || *response.Code < http.StatusBadRequest {
		return nil
	}

	body := ""
	if response.Body != nil {
		body = *response.Body
	}

	var payload struct {
		Error *ErrorResponse `json:"error"`
	}
	if json.Unmarshal([]byte(body), &payload) == nil && payload.Error != nil {
		return &APIError{
			StatusCode: *response.Code,
			ErrorInfo:  *payload.Error,
			RawBody:    body,
		}
	}
	return &APIError{StatusCode: *response.Code, RawBody: body}
}

func decodeBatchResponses(calls []BatchCall, responses []*BatchResponse) []*BatchCallError {
	errors := []*BatchCallError{}
	for index, call := range calls {
		if call.decoder == nil {
			continue
		}
		if index >= len(responses) {
			errors = append(errors, &BatchCallError{
				Index: index,
				Name:  call.Name,
				Err:   fmt.Errorf("facebook batch response is missing"),
			})
			continue
		}
		if err := call.decoder(responses[index]); err != nil {
			errors = append(errors, &BatchCallError{
				Index: index,
				Name:  call.Name,
				Err:   err,
			})
		}
	}
	return errors
}

func batchRequestParams(batchKey string, payload []batchCallPayload, options ...BatchExecuteOption) Params {
	config := newBatchExecuteOptions(options...)
	out := Params{}
	for key, value := range config.params {
		out[key] = value
	}
	out[batchKey] = payload
	return out
}

func newBatchExecuteOptions(options ...BatchExecuteOption) batchExecuteOptions {
	config := batchExecuteOptions{params: Params{}}
	for _, option := range options {
		if option != nil {
			option(&config)
		}
	}
	return config
}

func (options *batchExecuteOptions) ensureParams() {
	if options.params == nil {
		options.params = Params{}
	}
}

func batchPayloadForCalls(calls []BatchCall) ([]batchCallPayload, error) {
	payload := make([]batchCallPayload, 0, len(calls))
	for _, call := range calls {
		item, err := batchPayloadForCall(call)
		if err != nil {
			return nil, err
		}
		payload = append(payload, item)
	}
	return payload, nil
}

type batchCallPayload struct {
	Method                string        `json:"method"`
	RelativeURL           string        `json:"relative_url"`
	Body                  string        `json:"body,omitempty"`
	Name                  string        `json:"name,omitempty"`
	DependsOn             string        `json:"depends_on,omitempty"`
	OmitResponseOnSuccess *bool         `json:"omit_response_on_success,omitempty"`
	Headers               []BatchHeader `json:"headers,omitempty"`
}

func batchPayloadForCall(call BatchCall) (batchCallPayload, error) {
	method := strings.ToUpper(strings.TrimSpace(call.Method))
	if method == "" {
		method = http.MethodGet
	}

	item := batchCallPayload{
		Method:                method,
		RelativeURL:           cleanBatchRelativeURL(call.RelativeURL),
		Body:                  call.Body,
		Name:                  call.Name,
		DependsOn:             call.DependsOn,
		OmitResponseOnSuccess: call.OmitResponseOnSuccess,
		Headers:               call.Headers,
	}
	if item.RelativeURL == "" {
		return item, fmt.Errorf("facebook batch call relative url is empty")
	}

	values, files, err := encodeParams(call.Params)
	if err != nil {
		return item, err
	}
	if len(files) > 0 {
		return item, fmt.Errorf("facebook batch call file params are not supported")
	}
	applyBatchParams(&item, values)
	return item, nil
}

func cleanBatchRelativeURL(value string) string {
	return strings.TrimLeft(strings.TrimSpace(value), "/")
}

func applyBatchParams(item *batchCallPayload, values url.Values) {
	if len(values) == 0 {
		return
	}

	encoded := values.Encode()
	if item.Method == http.MethodGet || item.Method == http.MethodDelete {
		item.RelativeURL = appendBatchQuery(item.RelativeURL, encoded)
		return
	}

	if item.Body == "" {
		item.Body = encoded
		return
	}
	item.Body += "&" + encoded
}

func appendBatchQuery(relativeURL string, query string) string {
	if strings.Contains(relativeURL, "?") {
		return relativeURL + "&" + query
	}
	return relativeURL + "?" + query
}
