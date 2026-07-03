package objects

type AdCreativeDegreesOfFreedomSpec struct {
	AdHandleType                 *string                 `json:"ad_handle_type,omitempty"`
	CreativeFeaturesSpec         *AdCreativeFeaturesSpec `json:"creative_features_spec,omitempty"`
	DegreesOfFreedomType         *string                 `json:"degrees_of_freedom_type,omitempty"`
	ImageTransformationTypes     *[]string               `json:"image_transformation_types,omitempty"`
	MultiMediaTransformationType *string                 `json:"multi_media_transformation_type,omitempty"`
	StoriesTransformationTypes   *[]string               `json:"stories_transformation_types,omitempty"`
	TextTransformationTypes      *[]string               `json:"text_transformation_types,omitempty"`
	VideoTransformationTypes     *[]string               `json:"video_transformation_types,omitempty"`
}
