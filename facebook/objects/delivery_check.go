package objects

type DeliveryCheck struct {
	CheckName   *string                 `json:"check_name,omitempty"`
	Description *string                 `json:"description,omitempty"`
	ExtraInfo   *DeliveryCheckExtraInfo `json:"extra_info,omitempty"`
	Summary     *string                 `json:"summary,omitempty"`
}

var DeliveryCheckFields = struct {
	CheckName   string
	Description string
	ExtraInfo   string
	Summary     string
}{
	CheckName:   "check_name",
	Description: "description",
	ExtraInfo:   "extra_info",
	Summary:     "summary",
}
