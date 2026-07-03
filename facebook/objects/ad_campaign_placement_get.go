package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdCampaignPlacementGet struct {
	EffectiveAudienceNetworkPositions *[]enums.AdcampaignplacementgetEffectiveAudienceNetworkPositions `json:"effective_audience_network_positions,omitempty"`
	EffectiveDevicePlatforms          *[]enums.AdcampaignplacementgetEffectiveDevicePlatforms          `json:"effective_device_platforms,omitempty"`
	EffectiveFacebookPositions        *[]enums.AdcampaignplacementgetEffectiveFacebookPositions        `json:"effective_facebook_positions,omitempty"`
	EffectiveInstagramPositions       *[]enums.AdcampaignplacementgetEffectiveInstagramPositions       `json:"effective_instagram_positions,omitempty"`
	EffectiveMessengerPositions       *[]enums.AdcampaignplacementgetEffectiveMessengerPositions       `json:"effective_messenger_positions,omitempty"`
	EffectiveOculusPositions          *[]enums.AdcampaignplacementgetEffectiveOculusPositions          `json:"effective_oculus_positions,omitempty"`
	EffectivePublisherPlatforms       *[]enums.AdcampaignplacementgetEffectivePublisherPlatforms       `json:"effective_publisher_platforms,omitempty"`
	Metadata                          *map[string]interface{}                                          `json:"metadata,omitempty"`
	Recommendations                   *[]map[string]interface{}                                        `json:"recommendations,omitempty"`
}

var AdCampaignPlacementGetFields = struct {
	EffectiveAudienceNetworkPositions string
	EffectiveDevicePlatforms          string
	EffectiveFacebookPositions        string
	EffectiveInstagramPositions       string
	EffectiveMessengerPositions       string
	EffectiveOculusPositions          string
	EffectivePublisherPlatforms       string
	Metadata                          string
	Recommendations                   string
}{
	EffectiveAudienceNetworkPositions: "effective_audience_network_positions",
	EffectiveDevicePlatforms:          "effective_device_platforms",
	EffectiveFacebookPositions:        "effective_facebook_positions",
	EffectiveInstagramPositions:       "effective_instagram_positions",
	EffectiveMessengerPositions:       "effective_messenger_positions",
	EffectiveOculusPositions:          "effective_oculus_positions",
	EffectivePublisherPlatforms:       "effective_publisher_platforms",
	Metadata:                          "metadata",
	Recommendations:                   "recommendations",
}
