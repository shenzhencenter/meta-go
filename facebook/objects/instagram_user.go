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

var InstagramUserFields = struct {
	FollowCount        string
	FollowedByCount    string
	HasProfilePicture  string
	ID                 string
	IgUserID           string
	IsPrivate          string
	IsPublished        string
	MediaCount         string
	MiniShopStorefront string
	OwnerBusiness      string
	ProfilePic         string
	Username           string
}{
	FollowCount:        "follow_count",
	FollowedByCount:    "followed_by_count",
	HasProfilePicture:  "has_profile_picture",
	ID:                 "id",
	IgUserID:           "ig_user_id",
	IsPrivate:          "is_private",
	IsPublished:        "is_published",
	MediaCount:         "media_count",
	MiniShopStorefront: "mini_shop_storefront",
	OwnerBusiness:      "owner_business",
	ProfilePic:         "profile_pic",
	Username:           "username",
}
