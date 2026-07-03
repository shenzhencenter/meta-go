package enums

type AdaccountgeneratepreviewsCreativeFeatureEnumParam string

const (
	AdaccountgeneratepreviewsCreativeFeatureEnumParamIgVideoNativeSubtitle       AdaccountgeneratepreviewsCreativeFeatureEnumParam = "ig_video_native_subtitle"
	AdaccountgeneratepreviewsCreativeFeatureEnumParamImageAnimation              AdaccountgeneratepreviewsCreativeFeatureEnumParam = "image_animation"
	AdaccountgeneratepreviewsCreativeFeatureEnumParamProductBrowsing             AdaccountgeneratepreviewsCreativeFeatureEnumParam = "product_browsing"
	AdaccountgeneratepreviewsCreativeFeatureEnumParamProductMetadataAutomation   AdaccountgeneratepreviewsCreativeFeatureEnumParam = "product_metadata_automation"
	AdaccountgeneratepreviewsCreativeFeatureEnumParamProfileCard                 AdaccountgeneratepreviewsCreativeFeatureEnumParam = "profile_card"
	AdaccountgeneratepreviewsCreativeFeatureEnumParamStandardEnhancementsCatalog AdaccountgeneratepreviewsCreativeFeatureEnumParam = "standard_enhancements_catalog"
	AdaccountgeneratepreviewsCreativeFeatureEnumParamTextOverlayTranslation      AdaccountgeneratepreviewsCreativeFeatureEnumParam = "text_overlay_translation"
)

func (value AdaccountgeneratepreviewsCreativeFeatureEnumParam) String() string {
	return string(value)
}
