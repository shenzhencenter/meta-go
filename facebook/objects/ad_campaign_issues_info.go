package objects

type AdCampaignIssuesInfo struct {
	ErrorCode    *int    `json:"error_code,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	ErrorSummary *string `json:"error_summary,omitempty"`
	ErrorType    *string `json:"error_type,omitempty"`
	Level        *string `json:"level,omitempty"`
	Mid          *string `json:"mid,omitempty"`
}

var AdCampaignIssuesInfoFields = struct {
	ErrorCode    string
	ErrorMessage string
	ErrorSummary string
	ErrorType    string
	Level        string
	Mid          string
}{
	ErrorCode:    "error_code",
	ErrorMessage: "error_message",
	ErrorSummary: "error_summary",
	ErrorType:    "error_type",
	Level:        "level",
	Mid:          "mid",
}
