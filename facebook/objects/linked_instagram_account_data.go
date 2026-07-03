package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LinkedInstagramAccountData struct {
	AccessToken       *string  `json:"access_token,omitempty"`
	AnalyticsClaim    *string  `json:"analytics_claim,omitempty"`
	FullName          *string  `json:"full_name,omitempty"`
	ProfilePictureURL *string  `json:"profile_picture_url,omitempty"`
	UserID            *core.ID `json:"user_id,omitempty"`
	UserName          *string  `json:"user_name,omitempty"`
}

var LinkedInstagramAccountDataFields = struct {
	AccessToken       string
	AnalyticsClaim    string
	FullName          string
	ProfilePictureURL string
	UserID            string
	UserName          string
}{
	AccessToken:       "access_token",
	AnalyticsClaim:    "analytics_claim",
	FullName:          "full_name",
	ProfilePictureURL: "profile_picture_url",
	UserID:            "user_id",
	UserName:          "user_name",
}
