package objects

type BrandSafetyCampaignConfig struct {
	CommentModerationFilter *string `json:"comment_moderation_filter,omitempty"`
}

var BrandSafetyCampaignConfigFields = struct {
	CommentModerationFilter string
}{
	CommentModerationFilter: "comment_moderation_filter",
}
