package objects

type VideoCopyrightConditionGroup struct {
	Action         *string                   `json:"action,omitempty"`
	Conditions     *[]map[string]interface{} `json:"conditions,omitempty"`
	ValidityStatus *string                   `json:"validity_status,omitempty"`
}

var VideoCopyrightConditionGroupFields = struct {
	Action         string
	Conditions     string
	ValidityStatus string
}{
	Action:         "action",
	Conditions:     "conditions",
	ValidityStatus: "validity_status",
}
