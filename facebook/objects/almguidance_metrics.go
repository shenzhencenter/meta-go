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

var ALMGuidanceMetricsFields = struct {
	AdAccountID        string
	AdoptedObjects     string
	GuidanceName       string
	GuidanceType       string
	L28Adoption        string
	L28Available       string
	L28Click           string
	L28Conversion      string
	L28HasClick        string
	L28HasImpression   string
	L28Impression      string
	L28IsActioned      string
	L28IsAdopted       string
	L28IsAvailable     string
	L28IsPitched       string
	L28Pitch           string
	L28dAdoptedRevenue string
	LastActionedDs     string
	LastAdoptedDs      string
	LastPitchDs        string
	ParentAdvertiserID string
	ReportDs           string
}{
	AdAccountID:        "ad_account_id",
	AdoptedObjects:     "adopted_objects",
	GuidanceName:       "guidance_name",
	GuidanceType:       "guidance_type",
	L28Adoption:        "l28_adoption",
	L28Available:       "l28_available",
	L28Click:           "l28_click",
	L28Conversion:      "l28_conversion",
	L28HasClick:        "l28_has_click",
	L28HasImpression:   "l28_has_impression",
	L28Impression:      "l28_impression",
	L28IsActioned:      "l28_is_actioned",
	L28IsAdopted:       "l28_is_adopted",
	L28IsAvailable:     "l28_is_available",
	L28IsPitched:       "l28_is_pitched",
	L28Pitch:           "l28_pitch",
	L28dAdoptedRevenue: "l28d_adopted_revenue",
	LastActionedDs:     "last_actioned_ds",
	LastAdoptedDs:      "last_adopted_ds",
	LastPitchDs:        "last_pitch_ds",
	ParentAdvertiserID: "parent_advertiser_id",
	ReportDs:           "report_ds",
}
