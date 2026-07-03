package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BrandedContentAdError struct {
	BlameFieldSpec   *[]string `json:"blame_field_spec,omitempty"`
	ErrorCode        *int      `json:"error_code,omitempty"`
	ErrorDescription *string   `json:"error_description,omitempty"`
	ErrorMessage     *string   `json:"error_message,omitempty"`
	ErrorPlacement   *string   `json:"error_placement,omitempty"`
	ErrorSeverity    *string   `json:"error_severity,omitempty"`
	HelpCenterID     *core.ID  `json:"help_center_id,omitempty"`
}

var BrandedContentAdErrorFields = struct {
	BlameFieldSpec   string
	ErrorCode        string
	ErrorDescription string
	ErrorMessage     string
	ErrorPlacement   string
	ErrorSeverity    string
	HelpCenterID     string
}{
	BlameFieldSpec:   "blame_field_spec",
	ErrorCode:        "error_code",
	ErrorDescription: "error_description",
	ErrorMessage:     "error_message",
	ErrorPlacement:   "error_placement",
	ErrorSeverity:    "error_severity",
	HelpCenterID:     "help_center_id",
}
