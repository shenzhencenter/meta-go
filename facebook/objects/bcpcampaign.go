package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var BCPCampaignFields = struct {
	AdsPermissionRequired          string
	ApplicationDeadline            string
	CampaignGoal                   string
	CampaignGoalOther              string
	ContentDeliveryDeadline        string
	ContentDeliveryStartDate       string
	ContentRequirements            string
	ContentRequirementsDescription string
	Currency                       string
	DealNegotiationType            string
	Description                    string
	HasFreeProduct                 string
	ID                             string
	Name                           string
	PaymentAmountForAds            string
	PaymentAmountForContent        string
	PaymentDescription             string
}{
	AdsPermissionRequired:          "ads_permission_required",
	ApplicationDeadline:            "application_deadline",
	CampaignGoal:                   "campaign_goal",
	CampaignGoalOther:              "campaign_goal_other",
	ContentDeliveryDeadline:        "content_delivery_deadline",
	ContentDeliveryStartDate:       "content_delivery_start_date",
	ContentRequirements:            "content_requirements",
	ContentRequirementsDescription: "content_requirements_description",
	Currency:                       "currency",
	DealNegotiationType:            "deal_negotiation_type",
	Description:                    "description",
	HasFreeProduct:                 "has_free_product",
	ID:                             "id",
	Name:                           "name",
	PaymentAmountForAds:            "payment_amount_for_ads",
	PaymentAmountForContent:        "payment_amount_for_content",
	PaymentDescription:             "payment_description",
}
