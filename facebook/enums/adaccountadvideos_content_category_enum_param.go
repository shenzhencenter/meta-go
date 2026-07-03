package enums

type AdaccountadvideosContentCategoryEnumParam string

const (
	AdaccountadvideosContentCategoryEnumParamBeautyFashion AdaccountadvideosContentCategoryEnumParam = "BEAUTY_FASHION"
	AdaccountadvideosContentCategoryEnumParamBusiness      AdaccountadvideosContentCategoryEnumParam = "BUSINESS"
	AdaccountadvideosContentCategoryEnumParamCarsTrucks    AdaccountadvideosContentCategoryEnumParam = "CARS_TRUCKS"
	AdaccountadvideosContentCategoryEnumParamComedy        AdaccountadvideosContentCategoryEnumParam = "COMEDY"
	AdaccountadvideosContentCategoryEnumParamCuteAnimals   AdaccountadvideosContentCategoryEnumParam = "CUTE_ANIMALS"
	AdaccountadvideosContentCategoryEnumParamEntertainment AdaccountadvideosContentCategoryEnumParam = "ENTERTAINMENT"
	AdaccountadvideosContentCategoryEnumParamFamily        AdaccountadvideosContentCategoryEnumParam = "FAMILY"
	AdaccountadvideosContentCategoryEnumParamFoodHealth    AdaccountadvideosContentCategoryEnumParam = "FOOD_HEALTH"
	AdaccountadvideosContentCategoryEnumParamHome          AdaccountadvideosContentCategoryEnumParam = "HOME"
	AdaccountadvideosContentCategoryEnumParamLifestyle     AdaccountadvideosContentCategoryEnumParam = "LIFESTYLE"
	AdaccountadvideosContentCategoryEnumParamMusic         AdaccountadvideosContentCategoryEnumParam = "MUSIC"
	AdaccountadvideosContentCategoryEnumParamNews          AdaccountadvideosContentCategoryEnumParam = "NEWS"
	AdaccountadvideosContentCategoryEnumParamOther         AdaccountadvideosContentCategoryEnumParam = "OTHER"
	AdaccountadvideosContentCategoryEnumParamPolitics      AdaccountadvideosContentCategoryEnumParam = "POLITICS"
	AdaccountadvideosContentCategoryEnumParamScience       AdaccountadvideosContentCategoryEnumParam = "SCIENCE"
	AdaccountadvideosContentCategoryEnumParamSports        AdaccountadvideosContentCategoryEnumParam = "SPORTS"
	AdaccountadvideosContentCategoryEnumParamTechnology    AdaccountadvideosContentCategoryEnumParam = "TECHNOLOGY"
	AdaccountadvideosContentCategoryEnumParamVideoGaming   AdaccountadvideosContentCategoryEnumParam = "VIDEO_GAMING"
)

func (value AdaccountadvideosContentCategoryEnumParam) String() string {
	return string(value)
}
