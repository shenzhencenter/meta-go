package objects

type VideoCopyrightGeoGate struct {
	ExcludedCountries *[]string `json:"excluded_countries,omitempty"`
	IncludedCountries *[]string `json:"included_countries,omitempty"`
}

var VideoCopyrightGeoGateFields = struct {
	ExcludedCountries string
	IncludedCountries string
}{
	ExcludedCountries: "excluded_countries",
	IncludedCountries: "included_countries",
}
