package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type Placement struct {
	AudienceNetworkPositions          *[]string                                  `json:"audience_network_positions,omitempty"`
	DevicePlatforms                   *[]enums.PlacementDevicePlatforms          `json:"device_platforms,omitempty"`
	EffectiveAudienceNetworkPositions *[]string                                  `json:"effective_audience_network_positions,omitempty"`
	EffectiveDevicePlatforms          *[]enums.PlacementEffectiveDevicePlatforms `json:"effective_device_platforms,omitempty"`
	EffectiveFacebookPositions        *[]string                                  `json:"effective_facebook_positions,omitempty"`
	EffectiveInstagramPositions       *[]string                                  `json:"effective_instagram_positions,omitempty"`
	EffectiveMessengerPositions       *[]string                                  `json:"effective_messenger_positions,omitempty"`
	EffectiveOculusPositions          *[]string                                  `json:"effective_oculus_positions,omitempty"`
	EffectivePublisherPlatforms       *[]string                                  `json:"effective_publisher_platforms,omitempty"`
	EffectiveThreadsPositions         *[]string                                  `json:"effective_threads_positions,omitempty"`
	EffectiveWhatsappPositions        *[]string                                  `json:"effective_whatsapp_positions,omitempty"`
	FacebookPositions                 *[]string                                  `json:"facebook_positions,omitempty"`
	InstagramPositions                *[]string                                  `json:"instagram_positions,omitempty"`
	MessengerPositions                *[]string                                  `json:"messenger_positions,omitempty"`
	OculusPositions                   *[]string                                  `json:"oculus_positions,omitempty"`
	PublisherPlatforms                *[]string                                  `json:"publisher_platforms,omitempty"`
	ThreadsPositions                  *[]string                                  `json:"threads_positions,omitempty"`
	WhatsappPositions                 *[]string                                  `json:"whatsapp_positions,omitempty"`
}
