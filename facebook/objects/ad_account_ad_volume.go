package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountAdVolume struct {
	ActorID                                          *core.ID                  `json:"actor_id,omitempty"`
	ActorName                                        *string                   `json:"actor_name,omitempty"`
	AdLimitScopeBusiness                             *Business                 `json:"ad_limit_scope_business,omitempty"`
	AdLimitScopeBusinessManagerID                    *core.ID                  `json:"ad_limit_scope_business_manager_id,omitempty"`
	AdLimitSetByPageAdmin                            *uint64                   `json:"ad_limit_set_by_page_admin,omitempty"`
	AdsRunningOrInReviewCount                        *uint64                   `json:"ads_running_or_in_review_count,omitempty"`
	AdsRunningOrInReviewCountSubjectToLimitSetByPage *uint64                   `json:"ads_running_or_in_review_count_subject_to_limit_set_by_page,omitempty"`
	CurrentAccountAdsRunningOrInReviewCount          *uint64                   `json:"current_account_ads_running_or_in_review_count,omitempty"`
	FutureLimitActivationDate                        *string                   `json:"future_limit_activation_date,omitempty"`
	FutureLimitOnAdsRunningOrInReview                *uint64                   `json:"future_limit_on_ads_running_or_in_review,omitempty"`
	LimitOnAdsRunningOrInReview                      *uint64                   `json:"limit_on_ads_running_or_in_review,omitempty"`
	Recommendations                                  *[]map[string]interface{} `json:"recommendations,omitempty"`
}

var AdAccountAdVolumeFields = struct {
	ActorID                                          string
	ActorName                                        string
	AdLimitScopeBusiness                             string
	AdLimitScopeBusinessManagerID                    string
	AdLimitSetByPageAdmin                            string
	AdsRunningOrInReviewCount                        string
	AdsRunningOrInReviewCountSubjectToLimitSetByPage string
	CurrentAccountAdsRunningOrInReviewCount          string
	FutureLimitActivationDate                        string
	FutureLimitOnAdsRunningOrInReview                string
	LimitOnAdsRunningOrInReview                      string
	Recommendations                                  string
}{
	ActorID:                       "actor_id",
	ActorName:                     "actor_name",
	AdLimitScopeBusiness:          "ad_limit_scope_business",
	AdLimitScopeBusinessManagerID: "ad_limit_scope_business_manager_id",
	AdLimitSetByPageAdmin:         "ad_limit_set_by_page_admin",
	AdsRunningOrInReviewCount:     "ads_running_or_in_review_count",
	AdsRunningOrInReviewCountSubjectToLimitSetByPage: "ads_running_or_in_review_count_subject_to_limit_set_by_page",
	CurrentAccountAdsRunningOrInReviewCount:          "current_account_ads_running_or_in_review_count",
	FutureLimitActivationDate:                        "future_limit_activation_date",
	FutureLimitOnAdsRunningOrInReview:                "future_limit_on_ads_running_or_in_review",
	LimitOnAdsRunningOrInReview:                      "limit_on_ads_running_or_in_review",
	Recommendations:                                  "recommendations",
}
