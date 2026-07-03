package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type OffsiteSignalContainerBusinessObject struct {
	Business                        *Business `json:"business,omitempty"`
	ID                              *core.ID  `json:"id,omitempty"`
	IsEligibleForSharingToAdAccount *bool     `json:"is_eligible_for_sharing_to_ad_account,omitempty"`
	IsEligibleForSharingToBusiness  *bool     `json:"is_eligible_for_sharing_to_business,omitempty"`
	IsUnavailable                   *bool     `json:"is_unavailable,omitempty"`
	Name                            *string   `json:"name,omitempty"`
	PrimaryContainerID              *core.ID  `json:"primary_container_id,omitempty"`
}
