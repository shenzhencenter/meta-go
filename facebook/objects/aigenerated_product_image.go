package objects

type AIGeneratedProductImage struct {
	FlaggedForManualReview *bool   `json:"flagged_for_manual_review,omitempty"`
	TransformedImageURL    *string `json:"transformed_image_url,omitempty"`
}

var AIGeneratedProductImageFields = struct {
	FlaggedForManualReview string
	TransformedImageURL    string
}{
	FlaggedForManualReview: "flagged_for_manual_review",
	TransformedImageURL:    "transformed_image_url",
}
