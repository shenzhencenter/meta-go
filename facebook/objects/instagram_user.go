package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type InstagramUser struct {
	FollowCount        *int      `json:"follow_count,omitempty"`
	FollowedByCount    *int      `json:"followed_by_count,omitempty"`
	HasProfilePicture  *bool     `json:"has_profile_picture,omitempty"`
	ID                 *core.ID  `json:"id,omitempty"`
	IgUserID           *core.ID  `json:"ig_user_id,omitempty"`
	IsPrivate          *bool     `json:"is_private,omitempty"`
	IsPublished        *bool     `json:"is_published,omitempty"`
	MediaCount         *int      `json:"media_count,omitempty"`
	MiniShopStorefront *Shop     `json:"mini_shop_storefront,omitempty"`
	OwnerBusiness      *Business `json:"owner_business,omitempty"`
	ProfilePic         *string   `json:"profile_pic,omitempty"`
	Username           *string   `json:"username,omitempty"`
}
