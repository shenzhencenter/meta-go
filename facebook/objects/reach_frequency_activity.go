package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ReachFrequencyActivity struct {
	AccountID        *core.ID `json:"account_id,omitempty"`
	CampaignActive   *bool    `json:"campaign_active,omitempty"`
	CampaignStarted  *bool    `json:"campaign_started,omitempty"`
	CreativeUploaded *bool    `json:"creative_uploaded,omitempty"`
	IoApproved       *bool    `json:"io_approved,omitempty"`
	SfLink           *string  `json:"sf_link,omitempty"`
}
