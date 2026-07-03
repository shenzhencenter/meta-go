package objects

type TargetingGeoLocationElectoralDistrict struct {
	Country           *string `json:"country,omitempty"`
	DeprecationCode   *string `json:"deprecation_code,omitempty"`
	ElectoralDistrict *string `json:"electoral_district,omitempty"`
	Key               *string `json:"key,omitempty"`
	Name              *string `json:"name,omitempty"`
}
