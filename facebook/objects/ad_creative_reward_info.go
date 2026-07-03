package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeRewardInfo struct {
	RewardOfferID   *core.ID `json:"reward_offer_id,omitempty"`
	RewardProgramID *core.ID `json:"reward_program_id,omitempty"`
}

var AdCreativeRewardInfoFields = struct {
	RewardOfferID   string
	RewardProgramID string
}{
	RewardOfferID:   "reward_offer_id",
	RewardProgramID: "reward_program_id",
}
