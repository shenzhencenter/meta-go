package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var OffsiteSignalContainerBusinessObjectFields = struct {
	Business                        string
	ID                              string
	IsEligibleForSharingToAdAccount string
	IsEligibleForSharingToBusiness  string
	IsUnavailable                   string
	Name                            string
	PrimaryContainerID              string
}{
	Business:                        "business",
	ID:                              "id",
	IsEligibleForSharingToAdAccount: "is_eligible_for_sharing_to_ad_account",
	IsEligibleForSharingToBusiness:  "is_eligible_for_sharing_to_business",
	IsUnavailable:                   "is_unavailable",
	Name:                            "name",
	PrimaryContainerID:              "primary_container_id",
}
