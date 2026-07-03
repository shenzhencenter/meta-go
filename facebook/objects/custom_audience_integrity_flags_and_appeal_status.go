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
