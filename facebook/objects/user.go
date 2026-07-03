package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type User struct {
	About                           *string                `json:"about,omitempty"`
	AgeRange                        *AgeRange              `json:"age_range,omitempty"`
	Birthday                        *string                `json:"birthday,omitempty"`
	ClientBusinessID                *core.ID               `json:"client_business_id,omitempty"`
	Community                       *Group                 `json:"community,omitempty"`
	Cover                           *UserCoverPhoto        `json:"cover,omitempty"`
	Currency                        *Currency              `json:"currency,omitempty"`
	Education                       *[]EducationExperience `json:"education,omitempty"`
	Email                           *string                `json:"email,omitempty"`
	FavoriteAthletes                *[]Experience          `json:"favorite_athletes,omitempty"`
	FavoriteTeams                   *[]Experience          `json:"favorite_teams,omitempty"`
	FirstName                       *string                `json:"first_name,omitempty"`
	Gender                          *string                `json:"gender,omitempty"`
	Hometown                        *Page                  `json:"hometown,omitempty"`
	ID                              *core.ID               `json:"id,omitempty"`
	InspirationalPeople             *[]Experience          `json:"inspirational_people,omitempty"`
	InstallType                     *string                `json:"install_type,omitempty"`
	Installed                       *bool                  `json:"installed,omitempty"`
	IsGuestUser                     *bool                  `json:"is_guest_user,omitempty"`
	IsWorkAccount                   *bool                  `json:"is_work_account,omitempty"`
	Languages                       *[]Experience          `json:"languages,omitempty"`
	LastName                        *string                `json:"last_name,omitempty"`
	Link                            *string                `json:"link,omitempty"`
	LocalNewsMegaphoneDismissStatus *bool                  `json:"local_news_megaphone_dismiss_status,omitempty"`
	LocalNewsSubscriptionStatus     *bool                  `json:"local_news_subscription_status,omitempty"`
	Locale                          *string                `json:"locale,omitempty"`
	Location                        *Page                  `json:"location,omitempty"`
	MeetingFor                      *[]string              `json:"meeting_for,omitempty"`
	MiddleName                      *string                `json:"middle_name,omitempty"`
	Name                            *string                `json:"name,omitempty"`
	NameFormat                      *string                `json:"name_format,omitempty"`
	PaymentPricepoints              *PaymentPricepoints    `json:"payment_pricepoints,omitempty"`
	Political                       *string                `json:"political,omitempty"`
	ProfilePic                      *string                `json:"profile_pic,omitempty"`
	Quotes                          *string                `json:"quotes,omitempty"`
	RelationshipStatus              *string                `json:"relationship_status,omitempty"`
	Religion                        *string                `json:"religion,omitempty"`
	SharedLoginUpgradeRequiredBy    *core.Time             `json:"shared_login_upgrade_required_by,omitempty"`
	ShortName                       *string                `json:"short_name,omitempty"`
	SignificantOther                *User                  `json:"significant_other,omitempty"`
	Sports                          *[]Experience          `json:"sports,omitempty"`
	SupportsDonateButtonInLiveVideo *bool                  `json:"supports_donate_button_in_live_video,omitempty"`
	ThirdPartyID                    *core.ID               `json:"third_party_id,omitempty"`
	Timezone                        *float64               `json:"timezone,omitempty"`
	TokenForBusiness                *string                `json:"token_for_business,omitempty"`
	UpdatedTime                     *core.Time             `json:"updated_time,omitempty"`
	Verified                        *bool                  `json:"verified,omitempty"`
	VideoUploadLimits               *VideoUploadLimits     `json:"video_upload_limits,omitempty"`
	Website                         *string                `json:"website,omitempty"`
}

var UserFields = struct {
	About                           string
	AgeRange                        string
	Birthday                        string
	ClientBusinessID                string
	Community                       string
	Cover                           string
	Currency                        string
	Education                       string
	Email                           string
	FavoriteAthletes                string
	FavoriteTeams                   string
	FirstName                       string
	Gender                          string
	Hometown                        string
	ID                              string
	InspirationalPeople             string
	InstallType                     string
	Installed                       string
	IsGuestUser                     string
	IsWorkAccount                   string
	Languages                       string
	LastName                        string
	Link                            string
	LocalNewsMegaphoneDismissStatus string
	LocalNewsSubscriptionStatus     string
	Locale                          string
	Location                        string
	MeetingFor                      string
	MiddleName                      string
	Name                            string
	NameFormat                      string
	PaymentPricepoints              string
	Political                       string
	ProfilePic                      string
	Quotes                          string
	RelationshipStatus              string
	Religion                        string
	SharedLoginUpgradeRequiredBy    string
	ShortName                       string
	SignificantOther                string
	Sports                          string
	SupportsDonateButtonInLiveVideo string
	ThirdPartyID                    string
	Timezone                        string
	TokenForBusiness                string
	UpdatedTime                     string
	Verified                        string
	VideoUploadLimits               string
	Website                         string
}{
	About:                           "about",
	AgeRange:                        "age_range",
	Birthday:                        "birthday",
	ClientBusinessID:                "client_business_id",
	Community:                       "community",
	Cover:                           "cover",
	Currency:                        "currency",
	Education:                       "education",
	Email:                           "email",
	FavoriteAthletes:                "favorite_athletes",
	FavoriteTeams:                   "favorite_teams",
	FirstName:                       "first_name",
	Gender:                          "gender",
	Hometown:                        "hometown",
	ID:                              "id",
	InspirationalPeople:             "inspirational_people",
	InstallType:                     "install_type",
	Installed:                       "installed",
	IsGuestUser:                     "is_guest_user",
	IsWorkAccount:                   "is_work_account",
	Languages:                       "languages",
	LastName:                        "last_name",
	Link:                            "link",
	LocalNewsMegaphoneDismissStatus: "local_news_megaphone_dismiss_status",
	LocalNewsSubscriptionStatus:     "local_news_subscription_status",
	Locale:                          "locale",
	Location:                        "location",
	MeetingFor:                      "meeting_for",
	MiddleName:                      "middle_name",
	Name:                            "name",
	NameFormat:                      "name_format",
	PaymentPricepoints:              "payment_pricepoints",
	Political:                       "political",
	ProfilePic:                      "profile_pic",
	Quotes:                          "quotes",
	RelationshipStatus:              "relationship_status",
	Religion:                        "religion",
	SharedLoginUpgradeRequiredBy:    "shared_login_upgrade_required_by",
	ShortName:                       "short_name",
	SignificantOther:                "significant_other",
	Sports:                          "sports",
	SupportsDonateButtonInLiveVideo: "supports_donate_button_in_live_video",
	ThirdPartyID:                    "third_party_id",
	Timezone:                        "timezone",
	TokenForBusiness:                "token_for_business",
	UpdatedTime:                     "updated_time",
	Verified:                        "verified",
	VideoUploadLimits:               "video_upload_limits",
	Website:                         "website",
}
