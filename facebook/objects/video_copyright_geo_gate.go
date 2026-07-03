package objects

type VideoCopyrightGeoGate struct {
	ExcludedCountries *[]string `json:"excluded_countries,omitempty"`
	IncludedCountries *[]string `json:"included_countries,omitempty"`
}
