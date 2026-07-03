package enums

type AdaccountadcreativesCategoryMediaSourceEnumParam string

const (
	AdaccountadcreativesCategoryMediaSourceEnumParamCategory          AdaccountadcreativesCategoryMediaSourceEnumParam = "CATEGORY"
	AdaccountadcreativesCategoryMediaSourceEnumParamMixed             AdaccountadcreativesCategoryMediaSourceEnumParam = "MIXED"
	AdaccountadcreativesCategoryMediaSourceEnumParamProductsCollage   AdaccountadcreativesCategoryMediaSourceEnumParam = "PRODUCTS_COLLAGE"
	AdaccountadcreativesCategoryMediaSourceEnumParamProductsSlideshow AdaccountadcreativesCategoryMediaSourceEnumParam = "PRODUCTS_SLIDESHOW"
)

func (value AdaccountadcreativesCategoryMediaSourceEnumParam) String() string {
	return string(value)
}
