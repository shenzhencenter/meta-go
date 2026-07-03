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
