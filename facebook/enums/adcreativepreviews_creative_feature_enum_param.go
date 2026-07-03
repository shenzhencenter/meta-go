package enums

type AdcreativepreviewsCreativeFeatureEnumParam string

const (
	AdcreativepreviewsCreativeFeatureEnumParamIgVideoNativeSubtitle       AdcreativepreviewsCreativeFeatureEnumParam = "ig_video_native_subtitle"
	AdcreativepreviewsCreativeFeatureEnumParamImageAnimation              AdcreativepreviewsCreativeFeatureEnumParam = "image_animation"
	AdcreativepreviewsCreativeFeatureEnumParamProductBrowsing             AdcreativepreviewsCreativeFeatureEnumParam = "product_browsing"
	AdcreativepreviewsCreativeFeatureEnumParamProductMetadataAutomation   AdcreativepreviewsCreativeFeatureEnumParam = "product_metadata_automation"
	AdcreativepreviewsCreativeFeatureEnumParamProfileCard                 AdcreativepreviewsCreativeFeatureEnumParam = "profile_card"
	AdcreativepreviewsCreativeFeatureEnumParamStandardEnhancementsCatalog AdcreativepreviewsCreativeFeatureEnumParam = "standard_enhancements_catalog"
	AdcreativepreviewsCreativeFeatureEnumParamTextOverlayTranslation      AdcreativepreviewsCreativeFeatureEnumParam = "text_overlay_translation"
)

func (value AdcreativepreviewsCreativeFeatureEnumParam) String() string {
	return string(value)
}
