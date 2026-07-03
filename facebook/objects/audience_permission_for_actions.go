package objects

type AudiencePermissionForActions struct {
	CanEdit                    *bool `json:"can_edit,omitempty"`
	CanSeeInsight              *bool `json:"can_see_insight,omitempty"`
	CanShare                   *bool `json:"can_share,omitempty"`
	SubtypeSupportsLookalike   *bool `json:"subtype_supports_lookalike,omitempty"`
	SupportsRecipientLookalike *bool `json:"supports_recipient_lookalike,omitempty"`
}
