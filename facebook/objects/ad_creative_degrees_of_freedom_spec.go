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

var AdCreativeDegreesOfFreedomSpecFields = struct {
	AdHandleType                 string
	CreativeFeaturesSpec         string
	DegreesOfFreedomType         string
	ImageTransformationTypes     string
	MultiMediaTransformationType string
	StoriesTransformationTypes   string
	TextTransformationTypes      string
	VideoTransformationTypes     string
}{
	AdHandleType:                 "ad_handle_type",
	CreativeFeaturesSpec:         "creative_features_spec",
	DegreesOfFreedomType:         "degrees_of_freedom_type",
	ImageTransformationTypes:     "image_transformation_types",
	MultiMediaTransformationType: "multi_media_transformation_type",
	StoriesTransformationTypes:   "stories_transformation_types",
	TextTransformationTypes:      "text_transformation_types",
	VideoTransformationTypes:     "video_transformation_types",
}
