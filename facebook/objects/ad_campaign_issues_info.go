package objects

type AdCampaignIssuesInfo struct {
	ErrorCode    *int    `json:"error_code,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	ErrorSummary *string `json:"error_summary,omitempty"`
	ErrorType    *string `json:"error_type,omitempty"`
	Level        *string `json:"level,omitempty"`
	Mid          *string `json:"mid,omitempty"`
}
