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
