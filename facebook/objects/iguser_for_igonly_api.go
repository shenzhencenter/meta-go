package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type IGUserForIGOnlyAPI struct {
	AccountType       *string  `json:"account_type,omitempty"`
	Biography         *string  `json:"biography,omitempty"`
	FollowersCount    *int     `json:"followers_count,omitempty"`
	FollowsCount      *int     `json:"follows_count,omitempty"`
	ID                *core.ID `json:"id,omitempty"`
	MediaCount        *int     `json:"media_count,omitempty"`
	Name              *string  `json:"name,omitempty"`
	ProfilePictureURL *string  `json:"profile_picture_url,omitempty"`
	UserID            *core.ID `json:"user_id,omitempty"`
	Username          *string  `json:"username,omitempty"`
	Website           *string  `json:"website,omitempty"`
}

var IGUserForIGOnlyAPIFields = struct {
	AccountType       string
	Biography         string
	FollowersCount    string
	FollowsCount      string
	ID                string
	MediaCount        string
	Name              string
	ProfilePictureURL string
	UserID            string
	Username          string
	Website           string
}{
	AccountType:       "account_type",
	Biography:         "biography",
	FollowersCount:    "followers_count",
	FollowsCount:      "follows_count",
	ID:                "id",
	MediaCount:        "media_count",
	Name:              "name",
	ProfilePictureURL: "profile_picture_url",
	UserID:            "user_id",
	Username:          "username",
	Website:           "website",
}
