package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CPASLsbImageBank struct {
	AdGroupID             *core.ID `json:"ad_group_id,omitempty"`
	CatalogSegmentProxyID *core.ID `json:"catalog_segment_proxy_id,omitempty"`
	ID                    *core.ID `json:"id,omitempty"`
}

var CPASLsbImageBankFields = struct {
	AdGroupID             string
	CatalogSegmentProxyID string
	ID                    string
}{
	AdGroupID:             "ad_group_id",
	CatalogSegmentProxyID: "catalog_segment_proxy_id",
	ID:                    "id",
}
