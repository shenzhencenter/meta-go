package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
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

var AdAssetCustomizationRuleCustomizationSpecFields = struct {
	AgeMax                   string
	AgeMin                   string
	AudienceNetworkPositions string
	DevicePlatforms          string
	FacebookPositions        string
	GeoLocations             string
	InstagramPositions       string
	Locales                  string
	MessengerPositions       string
	PublisherPlatforms       string
	ThreadsPositions         string
}{
	AgeMax:                   "age_max",
	AgeMin:                   "age_min",
	AudienceNetworkPositions: "audience_network_positions",
	DevicePlatforms:          "device_platforms",
	FacebookPositions:        "facebook_positions",
	GeoLocations:             "geo_locations",
	InstagramPositions:       "instagram_positions",
	Locales:                  "locales",
	MessengerPositions:       "messenger_positions",
	PublisherPlatforms:       "publisher_platforms",
	ThreadsPositions:         "threads_positions",
}
