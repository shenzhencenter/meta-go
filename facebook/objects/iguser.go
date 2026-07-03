package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type IGUser struct {
	Biography                     *string                         `json:"biography,omitempty"`
	BusinessDiscovery             *IGUser                         `json:"business_discovery,omitempty"`
	CollaborativeMediaSearch      *ShadowIGUserCollaborativeMedia `json:"collaborative_media_search,omitempty"`
	FollowersCount                *int                            `json:"followers_count,omitempty"`
	FollowsCount                  *int                            `json:"follows_count,omitempty"`
	HasProfilePic                 *bool                           `json:"has_profile_pic,omitempty"`
	ID                            *core.ID                        `json:"id,omitempty"`
	IgID                          *core.ID                        `json:"ig_id,omitempty"`
	IsPublished                   *bool                           `json:"is_published,omitempty"`
	LegacyInstagramUserID         *core.ID                        `json:"legacy_instagram_user_id,omitempty"`
	MediaCount                    *int                            `json:"media_count,omitempty"`
	MentionedComment              *IGComment                      `json:"mentioned_comment,omitempty"`
	MentionedMedia                *IGMedia                        `json:"mentioned_media,omitempty"`
	Name                          *string                         `json:"name,omitempty"`
	OwnerBusiness                 *Business                       `json:"owner_business,omitempty"`
	ProfilePictureURL             *string                         `json:"profile_picture_url,omitempty"`
	ShoppingProductTagEligibility *bool                           `json:"shopping_product_tag_eligibility,omitempty"`
	ShoppingReviewStatus          *string                         `json:"shopping_review_status,omitempty"`
	Username                      *string                         `json:"username,omitempty"`
	Website                       *string                         `json:"website,omitempty"`
}

var IGUserFields = struct {
	Biography                     string
	BusinessDiscovery             string
	CollaborativeMediaSearch      string
	FollowersCount                string
	FollowsCount                  string
	HasProfilePic                 string
	ID                            string
	IgID                          string
	IsPublished                   string
	LegacyInstagramUserID         string
	MediaCount                    string
	MentionedComment              string
	MentionedMedia                string
	Name                          string
	OwnerBusiness                 string
	ProfilePictureURL             string
	ShoppingProductTagEligibility string
	ShoppingReviewStatus          string
	Username                      string
	Website                       string
}{
	Biography:                     "biography",
	BusinessDiscovery:             "business_discovery",
	CollaborativeMediaSearch:      "collaborative_media_search",
	FollowersCount:                "followers_count",
	FollowsCount:                  "follows_count",
	HasProfilePic:                 "has_profile_pic",
	ID:                            "id",
	IgID:                          "ig_id",
	IsPublished:                   "is_published",
	LegacyInstagramUserID:         "legacy_instagram_user_id",
	MediaCount:                    "media_count",
	MentionedComment:              "mentioned_comment",
	MentionedMedia:                "mentioned_media",
	Name:                          "name",
	OwnerBusiness:                 "owner_business",
	ProfilePictureURL:             "profile_picture_url",
	ShoppingProductTagEligibility: "shopping_product_tag_eligibility",
	ShoppingReviewStatus:          "shopping_review_status",
	Username:                      "username",
	Website:                       "website",
}
