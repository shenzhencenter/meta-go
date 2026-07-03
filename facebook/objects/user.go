package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
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
	SharedLoginUpgradeRequiredBy    *time.Time             `json:"shared_login_upgrade_required_by,omitempty"`
	ShortName                       *string                `json:"short_name,omitempty"`
	SignificantOther                *User                  `json:"significant_other,omitempty"`
	Sports                          *[]Experience          `json:"sports,omitempty"`
	SupportsDonateButtonInLiveVideo *bool                  `json:"supports_donate_button_in_live_video,omitempty"`
	ThirdPartyID                    *core.ID               `json:"third_party_id,omitempty"`
	Timezone                        *float64               `json:"timezone,omitempty"`
	TokenForBusiness                *string                `json:"token_for_business,omitempty"`
	UpdatedTime                     *time.Time             `json:"updated_time,omitempty"`
	Verified                        *bool                  `json:"verified,omitempty"`
	VideoUploadLimits               *VideoUploadLimits     `json:"video_upload_limits,omitempty"`
	Website                         *string                `json:"website,omitempty"`
}
