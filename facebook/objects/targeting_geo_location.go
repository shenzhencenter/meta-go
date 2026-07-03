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
