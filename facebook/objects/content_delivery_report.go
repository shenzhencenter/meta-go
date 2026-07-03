package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ContentDeliveryReport struct {
	ContentID            *core.ID `json:"content_id,omitempty"`
	CreatorID            *core.ID `json:"creator_id,omitempty"`
	EstimatedImpressions *uint64  `json:"estimated_impressions,omitempty"`
}
