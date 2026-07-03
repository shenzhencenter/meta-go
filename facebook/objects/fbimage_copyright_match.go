package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type FBImageCopyrightMatch struct {
	AddedToDashboardTime *core.Time                           `json:"added_to_dashboard_time,omitempty"`
	AppliedActions       *[]map[string]map[string]interface{} `json:"applied_actions,omitempty"`
	AuditLog             *[]map[string]interface{}            `json:"audit_log,omitempty"`
	AvailableUiActions   *[]string                            `json:"available_ui_actions,omitempty"`
	ExpirationDays       *int                                 `json:"expiration_days,omitempty"`
	GenericMatchData     *[]map[string]interface{}            `json:"generic_match_data,omitempty"`
	ID                   *core.ID                             `json:"id,omitempty"`
	IsBusinessPageMatch  *bool                                `json:"is_business_page_match,omitempty"`
	LastModifiedTime     *core.Time                           `json:"last_modified_time,omitempty"`
	MatchData            *[]map[string]interface{}            `json:"match_data,omitempty"`
	MatchStatus          *string                              `json:"match_status,omitempty"`
	OwnershipCountries   *VideoCopyrightGeoGate               `json:"ownership_countries,omitempty"`
	ReferenceOwner       *Profile                             `json:"reference_owner,omitempty"`
	TimeToAppeal         *int                                 `json:"time_to_appeal,omitempty"`
}

var FBImageCopyrightMatchFields = struct {
	AddedToDashboardTime string
	AppliedActions       string
	AuditLog             string
	AvailableUiActions   string
	ExpirationDays       string
	GenericMatchData     string
	ID                   string
	IsBusinessPageMatch  string
	LastModifiedTime     string
	MatchData            string
	MatchStatus          string
	OwnershipCountries   string
	ReferenceOwner       string
	TimeToAppeal         string
}{
	AddedToDashboardTime: "added_to_dashboard_time",
	AppliedActions:       "applied_actions",
	AuditLog:             "audit_log",
	AvailableUiActions:   "available_ui_actions",
	ExpirationDays:       "expiration_days",
	GenericMatchData:     "generic_match_data",
	ID:                   "id",
	IsBusinessPageMatch:  "is_business_page_match",
	LastModifiedTime:     "last_modified_time",
	MatchData:            "match_data",
	MatchStatus:          "match_status",
	OwnershipCountries:   "ownership_countries",
	ReferenceOwner:       "reference_owner",
	TimeToAppeal:         "time_to_appeal",
}
