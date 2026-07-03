package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PublisherBlockList struct {
	AppPublishers             *[]AppPublisher `json:"app_publishers,omitempty"`
	BusinessOwnerID           *core.ID        `json:"business_owner_id,omitempty"`
	ID                        *core.ID        `json:"id,omitempty"`
	IsAutoBlockingOn          *bool           `json:"is_auto_blocking_on,omitempty"`
	IsEligibleAtCampaignLevel *bool           `json:"is_eligible_at_campaign_level,omitempty"`
	LastUpdateTime            *core.Time      `json:"last_update_time,omitempty"`
	LastUpdateUser            *string         `json:"last_update_user,omitempty"`
	Name                      *string         `json:"name,omitempty"`
	OwnerAdAccountID          *core.ID        `json:"owner_ad_account_id,omitempty"`
	WebPublishers             *[]WebPublisher `json:"web_publishers,omitempty"`
}

var PublisherBlockListFields = struct {
	AppPublishers             string
	BusinessOwnerID           string
	ID                        string
	IsAutoBlockingOn          string
	IsEligibleAtCampaignLevel string
	LastUpdateTime            string
	LastUpdateUser            string
	Name                      string
	OwnerAdAccountID          string
	WebPublishers             string
}{
	AppPublishers:             "app_publishers",
	BusinessOwnerID:           "business_owner_id",
	ID:                        "id",
	IsAutoBlockingOn:          "is_auto_blocking_on",
	IsEligibleAtCampaignLevel: "is_eligible_at_campaign_level",
	LastUpdateTime:            "last_update_time",
	LastUpdateUser:            "last_update_user",
	Name:                      "name",
	OwnerAdAccountID:          "owner_ad_account_id",
	WebPublishers:             "web_publishers",
}
