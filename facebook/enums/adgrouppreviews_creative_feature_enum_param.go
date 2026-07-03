package enums

type AdgrouppreviewsCreativeFeatureEnumParam string

const (
	AdgrouppreviewsCreativeFeatureEnumParamIgVideoNativeSubtitle       AdgrouppreviewsCreativeFeatureEnumParam = "ig_video_native_subtitle"
	AdgrouppreviewsCreativeFeatureEnumParamImageAnimation              AdgrouppreviewsCreativeFeatureEnumParam = "image_animation"
	AdgrouppreviewsCreativeFeatureEnumParamProductBrowsing             AdgrouppreviewsCreativeFeatureEnumParam = "product_browsing"
	AdgrouppreviewsCreativeFeatureEnumParamProductMetadataAutomation   AdgrouppreviewsCreativeFeatureEnumParam = "product_metadata_automation"
	AdgrouppreviewsCreativeFeatureEnumParamProfileCard                 AdgrouppreviewsCreativeFeatureEnumParam = "profile_card"
	AdgrouppreviewsCreativeFeatureEnumParamStandardEnhancementsCatalog AdgrouppreviewsCreativeFeatureEnumParam = "standard_enhancements_catalog"
	AdgrouppreviewsCreativeFeatureEnumParamTextOverlayTranslation      AdgrouppreviewsCreativeFeatureEnumParam = "text_overlay_translation"
)

func (value AdgrouppreviewsCreativeFeatureEnumParam) String() string {
	return string(value)
}
