package objects

import (
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AdAssetCustomizationRuleCustomizationSpec struct {
	AgeMax                   *uint64                                                           `json:"age_max,omitempty"`
	AgeMin                   *uint64                                                           `json:"age_min,omitempty"`
	AudienceNetworkPositions *[]string                                                         `json:"audience_network_positions,omitempty"`
	DevicePlatforms          *[]enums.AdassetcustomizationrulecustomizationspecDevicePlatforms `json:"device_platforms,omitempty"`
	FacebookPositions        *[]string                                                         `json:"facebook_positions,omitempty"`
	GeoLocations             *TargetingGeoLocation                                             `json:"geo_locations,omitempty"`
	InstagramPositions       *[]string                                                         `json:"instagram_positions,omitempty"`
	Locales                  *[]uint64                                                         `json:"locales,omitempty"`
	MessengerPositions       *[]string                                                         `json:"messenger_positions,omitempty"`
	PublisherPlatforms       *[]string                                                         `json:"publisher_platforms,omitempty"`
	ThreadsPositions         *[]string                                                         `json:"threads_positions,omitempty"`
}
