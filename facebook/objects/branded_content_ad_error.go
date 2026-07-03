package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
