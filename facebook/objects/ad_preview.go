package objects

type AdPreview struct {
	Body               *string                 `json:"body,omitempty"`
	TransformationSpec *map[string]interface{} `json:"transformation_spec,omitempty"`
}
