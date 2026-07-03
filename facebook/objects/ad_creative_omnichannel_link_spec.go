package objects

type AdCreativeOmnichannelLinkSpec struct {
	App *map[string]interface{} `json:"app,omitempty"`
	Web *map[string]interface{} `json:"web,omitempty"`
}
