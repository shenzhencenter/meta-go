package objects

type TargetingGeoLocationPoliticalDistrict struct {
	Country           *string `json:"country,omitempty"`
	Key               *string `json:"key,omitempty"`
	Name              *string `json:"name,omitempty"`
	PoliticalDistrict *string `json:"political_district,omitempty"`
}

var TargetingGeoLocationPoliticalDistrictFields = struct {
	Country           string
	Key               string
	Name              string
	PoliticalDistrict string
}{
	Country:           "country",
	Key:               "key",
	Name:              "name",
	PoliticalDistrict: "political_district",
}
