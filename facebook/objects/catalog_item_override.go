package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CatalogItemOverride struct {
	ID                   *core.ID              `json:"id,omitempty"`
	LocalInfo            *ProductItemLocalInfo `json:"local_info,omitempty"`
	OverrideType         *string               `json:"override_type,omitempty"`
	OverrideValue        *string               `json:"override_value,omitempty"`
	UploadExpectedMethod *string               `json:"upload_expected_method,omitempty"`
}

var CatalogItemOverrideFields = struct {
	ID                   string
	LocalInfo            string
	OverrideType         string
	OverrideValue        string
	UploadExpectedMethod string
}{
	ID:                   "id",
	LocalInfo:            "local_info",
	OverrideType:         "override_type",
	OverrideValue:        "override_value",
	UploadExpectedMethod: "upload_expected_method",
}
