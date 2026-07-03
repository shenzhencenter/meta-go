package objects

type AudiencePermissionForActions struct {
	CanEdit                    *bool `json:"can_edit,omitempty"`
	CanSeeInsight              *bool `json:"can_see_insight,omitempty"`
	CanShare                   *bool `json:"can_share,omitempty"`
	SubtypeSupportsLookalike   *bool `json:"subtype_supports_lookalike,omitempty"`
	SupportsRecipientLookalike *bool `json:"supports_recipient_lookalike,omitempty"`
}

var AudiencePermissionForActionsFields = struct {
	CanEdit                    string
	CanSeeInsight              string
	CanShare                   string
	SubtypeSupportsLookalike   string
	SupportsRecipientLookalike string
}{
	CanEdit:                    "can_edit",
	CanSeeInsight:              "can_see_insight",
	CanShare:                   "can_share",
	SubtypeSupportsLookalike:   "subtype_supports_lookalike",
	SupportsRecipientLookalike: "supports_recipient_lookalike",
}
