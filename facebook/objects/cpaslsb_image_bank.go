package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CPASLsbImageBank struct {
	AdGroupID             *core.ID `json:"ad_group_id,omitempty"`
	CatalogSegmentProxyID *core.ID `json:"catalog_segment_proxy_id,omitempty"`
	ID                    *core.ID `json:"id,omitempty"`
}
