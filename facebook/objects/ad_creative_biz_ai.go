package objects

type AdCreativeBizAI struct {
	Capabilities    *map[string]interface{}   `json:"capabilities,omitempty"`
	Pills           *[]map[string]interface{} `json:"pills,omitempty"`
	PillsMbsVersion *int                      `json:"pills_mbs_version,omitempty"`
}
