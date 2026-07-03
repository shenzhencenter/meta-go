package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ALMGuidanceMetrics struct {
	AdAccountID        *core.ID                  `json:"ad_account_id,omitempty"`
	AdoptedObjects     *[]map[string]interface{} `json:"adopted_objects,omitempty"`
	GuidanceName       *string                   `json:"guidance_name,omitempty"`
	GuidanceType       *string                   `json:"guidance_type,omitempty"`
	L28Adoption        *int                      `json:"l28_adoption,omitempty"`
	L28Available       *int                      `json:"l28_available,omitempty"`
	L28Click           *int                      `json:"l28_click,omitempty"`
	L28Conversion      *int                      `json:"l28_conversion,omitempty"`
	L28HasClick        *bool                     `json:"l28_has_click,omitempty"`
	L28HasImpression   *bool                     `json:"l28_has_impression,omitempty"`
	L28Impression      *int                      `json:"l28_impression,omitempty"`
	L28IsActioned      *bool                     `json:"l28_is_actioned,omitempty"`
	L28IsAdopted       *bool                     `json:"l28_is_adopted,omitempty"`
	L28IsAvailable     *bool                     `json:"l28_is_available,omitempty"`
	L28IsPitched       *bool                     `json:"l28_is_pitched,omitempty"`
	L28Pitch           *int                      `json:"l28_pitch,omitempty"`
	L28dAdoptedRevenue *float64                  `json:"l28d_adopted_revenue,omitempty"`
	LastActionedDs     *string                   `json:"last_actioned_ds,omitempty"`
	LastAdoptedDs      *string                   `json:"last_adopted_ds,omitempty"`
	LastPitchDs        *string                   `json:"last_pitch_ds,omitempty"`
	ParentAdvertiserID *core.ID                  `json:"parent_advertiser_id,omitempty"`
	ReportDs           *string                   `json:"report_ds,omitempty"`
}
