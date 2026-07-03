package objects

type AdAssetFeedSpecCarousel struct {
	Adlabels            *[]AdAssetFeedSpecAssetLabel              `json:"adlabels,omitempty"`
	ChildAttachments    *[]AdAssetFeedSpecCarouselChildAttachment `json:"child_attachments,omitempty"`
	MultiShareEndCard   *bool                                     `json:"multi_share_end_card,omitempty"`
	MultiShareOptimized *bool                                     `json:"multi_share_optimized,omitempty"`
}

var AdAssetFeedSpecCarouselFields = struct {
	Adlabels            string
	ChildAttachments    string
	MultiShareEndCard   string
	MultiShareOptimized string
}{
	Adlabels:            "adlabels",
	ChildAttachments:    "child_attachments",
	MultiShareEndCard:   "multi_share_end_card",
	MultiShareOptimized: "multi_share_optimized",
}
