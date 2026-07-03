package objects

type AdPreview struct {
	Body               *string                 `json:"body,omitempty"`
	TransformationSpec *map[string]interface{} `json:"transformation_spec,omitempty"`
}

var AdPreviewFields = struct {
	Body               string
	TransformationSpec string
}{
	Body:               "body",
	TransformationSpec: "transformation_spec",
}
