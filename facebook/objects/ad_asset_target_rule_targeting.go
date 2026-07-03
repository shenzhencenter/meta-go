package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdAssetTargetRuleTargeting struct {
	AgeMax                   *uint64                                            `json:"age_max,omitempty"`
	AgeMin                   *uint64                                            `json:"age_min,omitempty"`
	AudienceNetworkPositions *[]string                                          `json:"audience_network_positions,omitempty"`
	DevicePlatforms          *[]enums.AdassettargetruletargetingDevicePlatforms `json:"device_platforms,omitempty"`
	FacebookPositions        *[]string                                          `json:"facebook_positions,omitempty"`
	GeoLocations             *TargetingGeoLocation                              `json:"geo_locations,omitempty"`
	InstagramPositions       *[]string                                          `json:"instagram_positions,omitempty"`
	PublisherPlatforms       *[]string                                          `json:"publisher_platforms,omitempty"`
	ThreadsPositions         *[]string                                          `json:"threads_positions,omitempty"`
	WhatsappPositions        *[]string                                          `json:"whatsapp_positions,omitempty"`
}
