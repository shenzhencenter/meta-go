package objects

type AdCreativeFeatureDetails struct {
	Customizations *AdCreativeFeatureCustomizations `json:"customizations,omitempty"`
	EnrollStatus   *string                          `json:"enroll_status,omitempty"`
}
