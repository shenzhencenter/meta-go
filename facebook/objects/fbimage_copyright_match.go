package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type FBImageCopyrightMatch struct {
	AddedToDashboardTime *time.Time                           `json:"added_to_dashboard_time,omitempty"`
	AppliedActions       *[]map[string]map[string]interface{} `json:"applied_actions,omitempty"`
	AuditLog             *[]map[string]interface{}            `json:"audit_log,omitempty"`
	AvailableUiActions   *[]string                            `json:"available_ui_actions,omitempty"`
	ExpirationDays       *int                                 `json:"expiration_days,omitempty"`
	GenericMatchData     *[]map[string]interface{}            `json:"generic_match_data,omitempty"`
	ID                   *core.ID                             `json:"id,omitempty"`
	IsBusinessPageMatch  *bool                                `json:"is_business_page_match,omitempty"`
	LastModifiedTime     *time.Time                           `json:"last_modified_time,omitempty"`
	MatchData            *[]map[string]interface{}            `json:"match_data,omitempty"`
	MatchStatus          *string                              `json:"match_status,omitempty"`
	OwnershipCountries   *VideoCopyrightGeoGate               `json:"ownership_countries,omitempty"`
	ReferenceOwner       *Profile                             `json:"reference_owner,omitempty"`
	TimeToAppeal         *int                                 `json:"time_to_appeal,omitempty"`
}
