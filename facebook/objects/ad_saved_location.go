package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdSavedLocation struct {
	Cities            *[]TargetingGeoLocationCity           `json:"cities,omitempty"`
	Countries         *[]string                             `json:"countries,omitempty"`
	CountryGroups     *[]string                             `json:"country_groups,omitempty"`
	CustomLocations   *[]TargetingGeoLocationCustomLocation `json:"custom_locations,omitempty"`
	GeoMarkets        *[]TargetingGeoLocationMarket         `json:"geo_markets,omitempty"`
	ID                *core.ID                              `json:"id,omitempty"`
	LocationSentences *[]string                             `json:"location_sentences,omitempty"`
	Name              *string                               `json:"name,omitempty"`
	Regions           *[]TargetingGeoLocationRegion         `json:"regions,omitempty"`
	Zips              *[]TargetingGeoLocationZip            `json:"zips,omitempty"`
}
