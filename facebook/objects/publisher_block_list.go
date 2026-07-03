package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type PublisherBlockList struct {
	AppPublishers             *[]AppPublisher `json:"app_publishers,omitempty"`
	BusinessOwnerID           *core.ID        `json:"business_owner_id,omitempty"`
	ID                        *core.ID        `json:"id,omitempty"`
	IsAutoBlockingOn          *bool           `json:"is_auto_blocking_on,omitempty"`
	IsEligibleAtCampaignLevel *bool           `json:"is_eligible_at_campaign_level,omitempty"`
	LastUpdateTime            *time.Time      `json:"last_update_time,omitempty"`
	LastUpdateUser            *string         `json:"last_update_user,omitempty"`
	Name                      *string         `json:"name,omitempty"`
	OwnerAdAccountID          *core.ID        `json:"owner_ad_account_id,omitempty"`
	WebPublishers             *[]WebPublisher `json:"web_publishers,omitempty"`
}
