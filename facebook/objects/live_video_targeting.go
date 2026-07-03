package objects

type LiveVideoTargeting struct {
	AgeMax            *uint64               `json:"age_max,omitempty"`
	AgeMin            *uint64               `json:"age_min,omitempty"`
	ExcludedCountries *[]string             `json:"excluded_countries,omitempty"`
	GeoLocations      *TargetingGeoLocation `json:"geo_locations,omitempty"`
}
