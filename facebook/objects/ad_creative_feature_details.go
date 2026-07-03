package objects

type AdCreativeFeatureDetails struct {
	Customizations *AdCreativeFeatureCustomizations `json:"customizations,omitempty"`
	EnrollStatus   *string                          `json:"enroll_status,omitempty"`
}

var AdCreativeFeatureDetailsFields = struct {
	Customizations string
	EnrollStatus   string
}{
	Customizations: "customizations",
	EnrollStatus:   "enroll_status",
}
