package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type IGUserExportForCAM struct {
	AgeBucket                        *string   `json:"age_bucket,omitempty"`
	Badges                           *[]string `json:"badges,omitempty"`
	Biography                        *string   `json:"biography,omitempty"`
	Country                          *string   `json:"country,omitempty"`
	Email                            *string   `json:"email,omitempty"`
	Gender                           *string   `json:"gender,omitempty"`
	HasBrandPartnershipExperience    *bool     `json:"has_brand_partnership_experience,omitempty"`
	ID                               *core.ID  `json:"id,omitempty"`
	IsAccountVerified                *bool     `json:"is_account_verified,omitempty"`
	IsCreatorFollowingBrand          *bool     `json:"is_creator_following_brand,omitempty"`
	IsPaidPartnershipMessagesEnabled *bool     `json:"is_paid_partnership_messages_enabled,omitempty"`
	MessagingID                      *core.ID  `json:"messaging_id,omitempty"`
	OnboardedStatus                  *bool     `json:"onboarded_status,omitempty"`
	PastBrandPartnershipPartners     *[]string `json:"past_brand_partnership_partners,omitempty"`
	PortfolioURL                     *string   `json:"portfolio_url,omitempty"`
	ProfilePictureURL                *string   `json:"profile_picture_url,omitempty"`
	Username                         *string   `json:"username,omitempty"`
}

var IGUserExportForCAMFields = struct {
	AgeBucket                        string
	Badges                           string
	Biography                        string
	Country                          string
	Email                            string
	Gender                           string
	HasBrandPartnershipExperience    string
	ID                               string
	IsAccountVerified                string
	IsCreatorFollowingBrand          string
	IsPaidPartnershipMessagesEnabled string
	MessagingID                      string
	OnboardedStatus                  string
	PastBrandPartnershipPartners     string
	PortfolioURL                     string
	ProfilePictureURL                string
	Username                         string
}{
	AgeBucket:                        "age_bucket",
	Badges:                           "badges",
	Biography:                        "biography",
	Country:                          "country",
	Email:                            "email",
	Gender:                           "gender",
	HasBrandPartnershipExperience:    "has_brand_partnership_experience",
	ID:                               "id",
	IsAccountVerified:                "is_account_verified",
	IsCreatorFollowingBrand:          "is_creator_following_brand",
	IsPaidPartnershipMessagesEnabled: "is_paid_partnership_messages_enabled",
	MessagingID:                      "messaging_id",
	OnboardedStatus:                  "onboarded_status",
	PastBrandPartnershipPartners:     "past_brand_partnership_partners",
	PortfolioURL:                     "portfolio_url",
	ProfilePictureURL:                "profile_picture_url",
	Username:                         "username",
}
