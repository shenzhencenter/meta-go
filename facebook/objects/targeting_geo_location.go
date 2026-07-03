package objects

type TargetingGeoLocation struct {
	Cities             *[]TargetingGeoLocationCity              `json:"cities,omitempty"`
	Countries          *[]string                                `json:"countries,omitempty"`
	CountryGroups      *[]string                                `json:"country_groups,omitempty"`
	CustomLocations    *[]TargetingGeoLocationCustomLocation    `json:"custom_locations,omitempty"`
	ElectoralDistricts *[]TargetingGeoLocationElectoralDistrict `json:"electoral_districts,omitempty"`
	GeoMarkets         *[]TargetingGeoLocationMarket            `json:"geo_markets,omitempty"`
	LargeGeoAreas      *[]TargetingGeoLocationGeoEntities       `json:"large_geo_areas,omitempty"`
	LocationClusterIds *[]TargetingGeoLocationLocationCluster   `json:"location_cluster_ids,omitempty"`
	LocationTypes      *[]string                                `json:"location_types,omitempty"`
	MediumGeoAreas     *[]TargetingGeoLocationGeoEntities       `json:"medium_geo_areas,omitempty"`
	MetroAreas         *[]TargetingGeoLocationGeoEntities       `json:"metro_areas,omitempty"`
	Neighborhoods      *[]TargetingGeoLocationGeoEntities       `json:"neighborhoods,omitempty"`
	Places             *[]TargetingGeoLocationPlace             `json:"places,omitempty"`
	PoliticalDistricts *[]TargetingGeoLocationPoliticalDistrict `json:"political_districts,omitempty"`
	Regions            *[]TargetingGeoLocationRegion            `json:"regions,omitempty"`
	SmallGeoAreas      *[]TargetingGeoLocationGeoEntities       `json:"small_geo_areas,omitempty"`
	Subcities          *[]TargetingGeoLocationGeoEntities       `json:"subcities,omitempty"`
	Subneighborhoods   *[]TargetingGeoLocationGeoEntities       `json:"subneighborhoods,omitempty"`
	Zips               *[]TargetingGeoLocationZip               `json:"zips,omitempty"`
}

var TargetingGeoLocationFields = struct {
	Cities             string
	Countries          string
	CountryGroups      string
	CustomLocations    string
	ElectoralDistricts string
	GeoMarkets         string
	LargeGeoAreas      string
	LocationClusterIds string
	LocationTypes      string
	MediumGeoAreas     string
	MetroAreas         string
	Neighborhoods      string
	Places             string
	PoliticalDistricts string
	Regions            string
	SmallGeoAreas      string
	Subcities          string
	Subneighborhoods   string
	Zips               string
}{
	Cities:             "cities",
	Countries:          "countries",
	CountryGroups:      "country_groups",
	CustomLocations:    "custom_locations",
	ElectoralDistricts: "electoral_districts",
	GeoMarkets:         "geo_markets",
	LargeGeoAreas:      "large_geo_areas",
	LocationClusterIds: "location_cluster_ids",
	LocationTypes:      "location_types",
	MediumGeoAreas:     "medium_geo_areas",
	MetroAreas:         "metro_areas",
	Neighborhoods:      "neighborhoods",
	Places:             "places",
	PoliticalDistricts: "political_districts",
	Regions:            "regions",
	SmallGeoAreas:      "small_geo_areas",
	Subcities:          "subcities",
	Subneighborhoods:   "subneighborhoods",
	Zips:               "zips",
}
