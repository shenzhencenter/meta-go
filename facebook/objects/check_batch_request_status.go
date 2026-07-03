package objects

type CheckBatchRequestStatus struct {
	Errors               *[]map[string]interface{} `json:"errors,omitempty"`
	ErrorsTotalCount     *int                      `json:"errors_total_count,omitempty"`
	Handle               *string                   `json:"handle,omitempty"`
	IdsOfInvalidRequests *[]string                 `json:"ids_of_invalid_requests,omitempty"`
	Status               *string                   `json:"status,omitempty"`
	Warnings             *[]map[string]interface{} `json:"warnings,omitempty"`
	WarningsTotalCount   *int                      `json:"warnings_total_count,omitempty"`
}

var CheckBatchRequestStatusFields = struct {
	Errors               string
	ErrorsTotalCount     string
	Handle               string
	IdsOfInvalidRequests string
	Status               string
	Warnings             string
	WarningsTotalCount   string
}{
	Errors:               "errors",
	ErrorsTotalCount:     "errors_total_count",
	Handle:               "handle",
	IdsOfInvalidRequests: "ids_of_invalid_requests",
	Status:               "status",
	Warnings:             "warnings",
	WarningsTotalCount:   "warnings_total_count",
}
