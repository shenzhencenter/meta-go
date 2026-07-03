package objects

type CustomAudienceIntegrityFlagsAndAppealStatus struct {
	CloseoutTime           *uint64   `json:"closeout_time,omitempty"`
	DaysUntilEnforcement   *uint64   `json:"days_until_enforcement,omitempty"`
	FlaggedFields          *[]string `json:"flagged_fields,omitempty"`
	IsEnforcementRolledOut *bool     `json:"is_enforcement_rolled_out,omitempty"`
	LatestAppealRequestor  *string   `json:"latest_appeal_requestor,omitempty"`
	LatestAppealTime       *uint64   `json:"latest_appeal_time,omitempty"`
	RestrictionStatus      *string   `json:"restriction_status,omitempty"`
}

var CustomAudienceIntegrityFlagsAndAppealStatusFields = struct {
	CloseoutTime           string
	DaysUntilEnforcement   string
	FlaggedFields          string
	IsEnforcementRolledOut string
	LatestAppealRequestor  string
	LatestAppealTime       string
	RestrictionStatus      string
}{
	CloseoutTime:           "closeout_time",
	DaysUntilEnforcement:   "days_until_enforcement",
	FlaggedFields:          "flagged_fields",
	IsEnforcementRolledOut: "is_enforcement_rolled_out",
	LatestAppealRequestor:  "latest_appeal_requestor",
	LatestAppealTime:       "latest_appeal_time",
	RestrictionStatus:      "restriction_status",
}
