package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
