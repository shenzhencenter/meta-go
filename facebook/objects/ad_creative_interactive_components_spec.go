package objects

type AdCreativeInteractiveComponentsSpec struct {
	ChildAttachments *[]map[string]interface{} `json:"child_attachments,omitempty"`
	Components       *[]map[string]interface{} `json:"components,omitempty"`
}
