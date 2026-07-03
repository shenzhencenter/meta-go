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
