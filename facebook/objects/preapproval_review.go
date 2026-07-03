package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PreapprovalReview struct {
	CompType        *string                              `json:"comp_type,omitempty"`
	CrowComponentID *core.ID                             `json:"crow_component_id,omitempty"`
	IsHumanReviewed *bool                                `json:"is_human_reviewed,omitempty"`
	IsReviewed      *bool                                `json:"is_reviewed,omitempty"`
	PolicyInfo      *[]map[string]map[string]interface{} `json:"policy_info,omitempty"`
}
