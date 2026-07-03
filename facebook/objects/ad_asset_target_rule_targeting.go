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

var AdAssetTargetRuleTargetingFields = struct {
	AgeMax                   string
	AgeMin                   string
	AudienceNetworkPositions string
	DevicePlatforms          string
	FacebookPositions        string
	GeoLocations             string
	InstagramPositions       string
	PublisherPlatforms       string
	ThreadsPositions         string
	WhatsappPositions        string
}{
	AgeMax:                   "age_max",
	AgeMin:                   "age_min",
	AudienceNetworkPositions: "audience_network_positions",
	DevicePlatforms:          "device_platforms",
	FacebookPositions:        "facebook_positions",
	GeoLocations:             "geo_locations",
	InstagramPositions:       "instagram_positions",
	PublisherPlatforms:       "publisher_platforms",
	ThreadsPositions:         "threads_positions",
	WhatsappPositions:        "whatsapp_positions",
}
