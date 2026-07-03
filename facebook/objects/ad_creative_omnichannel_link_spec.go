package objects

type AdCreativeOmnichannelLinkSpec struct {
	App *map[string]interface{} `json:"app,omitempty"`
	Web *map[string]interface{} `json:"web,omitempty"`
}

var AdCreativeOmnichannelLinkSpecFields = struct {
	App string
	Web string
}{
	App: "app",
	Web: "web",
}
