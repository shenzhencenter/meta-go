package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type BCPCampaign struct {
	AdsPermissionRequired          *bool                `json:"ads_permission_required,omitempty"`
	ApplicationDeadline            *string              `json:"application_deadline,omitempty"`
	CampaignGoal                   *string              `json:"campaign_goal,omitempty"`
	CampaignGoalOther              *string              `json:"campaign_goal_other,omitempty"`
	ContentDeliveryDeadline        *string              `json:"content_delivery_deadline,omitempty"`
	ContentDeliveryStartDate       *string              `json:"content_delivery_start_date,omitempty"`
	ContentRequirements            *[]map[string]uint64 `json:"content_requirements,omitempty"`
	ContentRequirementsDescription *string              `json:"content_requirements_description,omitempty"`
	Currency                       *string              `json:"currency,omitempty"`
	DealNegotiationType            *string              `json:"deal_negotiation_type,omitempty"`
	Description                    *string              `json:"description,omitempty"`
	HasFreeProduct                 *bool                `json:"has_free_product,omitempty"`
	ID                             *core.ID             `json:"id,omitempty"`
	Name                           *string              `json:"name,omitempty"`
	PaymentAmountForAds            *uint64              `json:"payment_amount_for_ads,omitempty"`
	PaymentAmountForContent        *uint64              `json:"payment_amount_for_content,omitempty"`
	PaymentDescription             *string              `json:"payment_description,omitempty"`
}
