package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdCreativeRewardInfo struct {
	RewardOfferID   *core.ID `json:"reward_offer_id,omitempty"`
	RewardProgramID *core.ID `json:"reward_program_id,omitempty"`
}
