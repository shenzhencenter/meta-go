package objects

type AdVolume struct {
	AdVolumeBreakDown                 *[]map[string]interface{} `json:"ad_volume_break_down,omitempty"`
	AdsRunningOrInReviewCount         *uint64                   `json:"ads_running_or_in_review_count,omitempty"`
	FutureLimitActivationDate         *string                   `json:"future_limit_activation_date,omitempty"`
	FutureLimitOnAdsRunningOrInReview *uint64                   `json:"future_limit_on_ads_running_or_in_review,omitempty"`
	IndividualAccountsAdVolume        *int                      `json:"individual_accounts_ad_volume,omitempty"`
	IsGpaPage                         *bool                     `json:"is_gpa_page,omitempty"`
	LimitOnAdsRunningOrInReview       *uint64                   `json:"limit_on_ads_running_or_in_review,omitempty"`
	OwningBusinessAdVolume            *int                      `json:"owning_business_ad_volume,omitempty"`
	PartnerBusinessAdVolume           *int                      `json:"partner_business_ad_volume,omitempty"`
	UserRole                          *string                   `json:"user_role,omitempty"`
}

var AdVolumeFields = struct {
	AdVolumeBreakDown                 string
	AdsRunningOrInReviewCount         string
	FutureLimitActivationDate         string
	FutureLimitOnAdsRunningOrInReview string
	IndividualAccountsAdVolume        string
	IsGpaPage                         string
	LimitOnAdsRunningOrInReview       string
	OwningBusinessAdVolume            string
	PartnerBusinessAdVolume           string
	UserRole                          string
}{
	AdVolumeBreakDown:                 "ad_volume_break_down",
	AdsRunningOrInReviewCount:         "ads_running_or_in_review_count",
	FutureLimitActivationDate:         "future_limit_activation_date",
	FutureLimitOnAdsRunningOrInReview: "future_limit_on_ads_running_or_in_review",
	IndividualAccountsAdVolume:        "individual_accounts_ad_volume",
	IsGpaPage:                         "is_gpa_page",
	LimitOnAdsRunningOrInReview:       "limit_on_ads_running_or_in_review",
	OwningBusinessAdVolume:            "owning_business_ad_volume",
	PartnerBusinessAdVolume:           "partner_business_ad_volume",
	UserRole:                          "user_role",
}
