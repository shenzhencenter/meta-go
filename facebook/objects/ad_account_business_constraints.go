package objects

import (
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AdAccountBusinessConstraints struct {
	AudienceControls        *map[string]interface{}                   `json:"audience_controls,omitempty"`
	CampaignsWithError      *[]string                                 `json:"campaigns_with_error,omitempty"`
	IsAgeRestrictionEnabled *bool                                     `json:"is_age_restriction_enabled,omitempty"`
	PlacementControls       *map[string]interface{}                   `json:"placement_controls,omitempty"`
	Status                  *enums.AdaccountbusinessconstraintsStatus `json:"status,omitempty"`
}
