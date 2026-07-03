package enums

type AdcreativelinkdataFormatOption string

const (
	AdcreativelinkdataFormatOptionCarouselArEffects        AdcreativelinkdataFormatOption = "carousel_ar_effects"
	AdcreativelinkdataFormatOptionCarouselImagesMultiItems AdcreativelinkdataFormatOption = "carousel_images_multi_items"
	AdcreativelinkdataFormatOptionCarouselImagesSingleItem AdcreativelinkdataFormatOption = "carousel_images_single_item"
	AdcreativelinkdataFormatOptionCarouselSlideshows       AdcreativelinkdataFormatOption = "carousel_slideshows"
	AdcreativelinkdataFormatOptionCollectionVideo          AdcreativelinkdataFormatOption = "collection_video"
	AdcreativelinkdataFormatOptionSingleImage              AdcreativelinkdataFormatOption = "single_image"
)

func (value AdcreativelinkdataFormatOption) String() string {
	return string(value)
}
