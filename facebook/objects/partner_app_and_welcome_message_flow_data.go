package objects

type PartnerAppAndWelcomeMessageFlowData struct {
	FlowData *[]map[string]interface{} `json:"flow_data,omitempty"`
}

var PartnerAppAndWelcomeMessageFlowDataFields = struct {
	FlowData string
}{
	FlowData: "flow_data",
}
