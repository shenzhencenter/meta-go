package objects

type AIGeneratedProductImage struct {
	FlaggedForManualReview *bool   `json:"flagged_for_manual_review,omitempty"`
	TransformedImageURL    *string `json:"transformed_image_url,omitempty"`
}
