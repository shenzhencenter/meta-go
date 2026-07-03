package objects

type TargetingGeoLocationLocationExpansion struct {
	Allowed *bool   `json:"allowed,omitempty"`
	Intent  *string `json:"intent,omitempty"`
}

var TargetingGeoLocationLocationExpansionFields = struct {
	Allowed string
	Intent  string
}{
	Allowed: "allowed",
	Intent:  "intent",
}
