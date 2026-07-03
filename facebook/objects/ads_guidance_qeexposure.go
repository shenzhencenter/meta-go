package objects

type AdsGuidanceQEExposure struct {
	AccountExposed *bool `json:"account_exposed,omitempty"`
}

var AdsGuidanceQEExposureFields = struct {
	AccountExposed string
}{
	AccountExposed: "account_exposed",
}
