package objects

type AdAccountURLForAssetExtraction struct {
	SourceType *string `json:"source_type,omitempty"`
	SourceURL  *string `json:"source_url,omitempty"`
}

var AdAccountURLForAssetExtractionFields = struct {
	SourceType string
	SourceURL  string
}{
	SourceType: "source_type",
	SourceURL:  "source_url",
}
