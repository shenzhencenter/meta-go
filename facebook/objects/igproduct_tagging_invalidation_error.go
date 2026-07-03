package objects

type IGProductTaggingInvalidationError struct {
	Description      *string `json:"description,omitempty"`
	TaggabilityState *string `json:"taggability_state,omitempty"`
	Title            *string `json:"title,omitempty"`
}

var IGProductTaggingInvalidationErrorFields = struct {
	Description      string
	TaggabilityState string
	Title            string
}{
	Description:      "description",
	TaggabilityState: "taggability_state",
	Title:            "title",
}
