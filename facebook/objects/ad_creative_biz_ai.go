package objects

type AdCreativeBizAI struct {
	Capabilities    *map[string]interface{}   `json:"capabilities,omitempty"`
	Pills           *[]map[string]interface{} `json:"pills,omitempty"`
	PillsMbsVersion *int                      `json:"pills_mbs_version,omitempty"`
}

var AdCreativeBizAIFields = struct {
	Capabilities    string
	Pills           string
	PillsMbsVersion string
}{
	Capabilities:    "capabilities",
	Pills:           "pills",
	PillsMbsVersion: "pills_mbs_version",
}
