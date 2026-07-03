package objects

type LiveVideoTargeting struct {
	AgeMax            *uint64               `json:"age_max,omitempty"`
	AgeMin            *uint64               `json:"age_min,omitempty"`
	ExcludedCountries *[]string             `json:"excluded_countries,omitempty"`
	GeoLocations      *TargetingGeoLocation `json:"geo_locations,omitempty"`
}

var LiveVideoTargetingFields = struct {
	AgeMax            string
	AgeMin            string
	ExcludedCountries string
	GeoLocations      string
}{
	AgeMax:            "age_max",
	AgeMin:            "age_min",
	ExcludedCountries: "excluded_countries",
	GeoLocations:      "geo_locations",
}
