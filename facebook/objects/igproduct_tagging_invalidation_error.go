package objects

type IGProductTaggingInvalidationError struct {
	Description      *string `json:"description,omitempty"`
	TaggabilityState *string `json:"taggability_state,omitempty"`
	Title            *string `json:"title,omitempty"`
}
