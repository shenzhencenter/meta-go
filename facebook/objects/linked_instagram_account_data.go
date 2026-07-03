package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type LinkedInstagramAccountData struct {
	AccessToken       *string  `json:"access_token,omitempty"`
	AnalyticsClaim    *string  `json:"analytics_claim,omitempty"`
	FullName          *string  `json:"full_name,omitempty"`
	ProfilePictureURL *string  `json:"profile_picture_url,omitempty"`
	UserID            *core.ID `json:"user_id,omitempty"`
	UserName          *string  `json:"user_name,omitempty"`
}
