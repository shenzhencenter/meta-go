package objects

type VideoCopyrightConditionGroup struct {
	Action         *string                   `json:"action,omitempty"`
	Conditions     *[]map[string]interface{} `json:"conditions,omitempty"`
	ValidityStatus *string                   `json:"validity_status,omitempty"`
}
