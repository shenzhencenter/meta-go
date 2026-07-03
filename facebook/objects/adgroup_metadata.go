package objects

type AdgroupMetadata struct {
	AdStandardEnhancementsEditSource *int    `json:"ad_standard_enhancements_edit_source,omitempty"`
	AdgroupCreationSource            *string `json:"adgroup_creation_source,omitempty"`
	AdgroupEditSource                *string `json:"adgroup_edit_source,omitempty"`
	AdgroupMediaSource               *string `json:"adgroup_media_source,omitempty"`
	CarouselStyle                    *string `json:"carousel_style,omitempty"`
	CarouselWithStaticCardStyle      *string `json:"carousel_with_static_card_style,omitempty"`
	IsPcaUnifiedFormatAd             *bool   `json:"is_pca_unified_format_ad,omitempty"`
}

var AdgroupMetadataFields = struct {
	AdStandardEnhancementsEditSource string
	AdgroupCreationSource            string
	AdgroupEditSource                string
	AdgroupMediaSource               string
	CarouselStyle                    string
	CarouselWithStaticCardStyle      string
	IsPcaUnifiedFormatAd             string
}{
	AdStandardEnhancementsEditSource: "ad_standard_enhancements_edit_source",
	AdgroupCreationSource:            "adgroup_creation_source",
	AdgroupEditSource:                "adgroup_edit_source",
	AdgroupMediaSource:               "adgroup_media_source",
	CarouselStyle:                    "carousel_style",
	CarouselWithStaticCardStyle:      "carousel_with_static_card_style",
	IsPcaUnifiedFormatAd:             "is_pca_unified_format_ad",
}
