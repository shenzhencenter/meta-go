package objects

type ProductSetMetadata struct {
	CoverImageURL         *string `json:"cover_image_url,omitempty"`
	Description           *string `json:"description,omitempty"`
	ExternalURL           *string `json:"external_url,omitempty"`
	IntegrityReviewStatus *string `json:"integrity_review_status,omitempty"`
}

var ProductSetMetadataFields = struct {
	CoverImageURL         string
	Description           string
	ExternalURL           string
	IntegrityReviewStatus string
}{
	CoverImageURL:         "cover_image_url",
	Description:           "description",
	ExternalURL:           "external_url",
	IntegrityReviewStatus: "integrity_review_status",
}
