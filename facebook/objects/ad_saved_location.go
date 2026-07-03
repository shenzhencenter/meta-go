package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var AdSavedLocationFields = struct {
	Cities            string
	Countries         string
	CountryGroups     string
	CustomLocations   string
	GeoMarkets        string
	ID                string
	LocationSentences string
	Name              string
	Regions           string
	Zips              string
}{
	Cities:            "cities",
	Countries:         "countries",
	CountryGroups:     "country_groups",
	CustomLocations:   "custom_locations",
	GeoMarkets:        "geo_markets",
	ID:                "id",
	LocationSentences: "location_sentences",
	Name:              "name",
	Regions:           "regions",
	Zips:              "zips",
}
