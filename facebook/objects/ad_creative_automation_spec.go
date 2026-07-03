package objects

type AdCreativeAutomationSpec struct {
	DecisionType     *string `json:"decision_type,omitempty"`
	EnrollmentStatus *string `json:"enrollment_status,omitempty"`
}

var AdCreativeAutomationSpecFields = struct {
	DecisionType     string
	EnrollmentStatus string
}{
	DecisionType:     "decision_type",
	EnrollmentStatus: "enrollment_status",
}
