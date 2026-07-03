package enums

type BusinessvideosContentCategoryEnumParam string

const (
	BusinessvideosContentCategoryEnumParamBeautyFashion BusinessvideosContentCategoryEnumParam = "BEAUTY_FASHION"
	BusinessvideosContentCategoryEnumParamBusiness      BusinessvideosContentCategoryEnumParam = "BUSINESS"
	BusinessvideosContentCategoryEnumParamCarsTrucks    BusinessvideosContentCategoryEnumParam = "CARS_TRUCKS"
	BusinessvideosContentCategoryEnumParamComedy        BusinessvideosContentCategoryEnumParam = "COMEDY"
	BusinessvideosContentCategoryEnumParamCuteAnimals   BusinessvideosContentCategoryEnumParam = "CUTE_ANIMALS"
	BusinessvideosContentCategoryEnumParamEntertainment BusinessvideosContentCategoryEnumParam = "ENTERTAINMENT"
	BusinessvideosContentCategoryEnumParamFamily        BusinessvideosContentCategoryEnumParam = "FAMILY"
	BusinessvideosContentCategoryEnumParamFoodHealth    BusinessvideosContentCategoryEnumParam = "FOOD_HEALTH"
	BusinessvideosContentCategoryEnumParamHome          BusinessvideosContentCategoryEnumParam = "HOME"
	BusinessvideosContentCategoryEnumParamLifestyle     BusinessvideosContentCategoryEnumParam = "LIFESTYLE"
	BusinessvideosContentCategoryEnumParamMusic         BusinessvideosContentCategoryEnumParam = "MUSIC"
	BusinessvideosContentCategoryEnumParamNews          BusinessvideosContentCategoryEnumParam = "NEWS"
	BusinessvideosContentCategoryEnumParamOther         BusinessvideosContentCategoryEnumParam = "OTHER"
	BusinessvideosContentCategoryEnumParamPolitics      BusinessvideosContentCategoryEnumParam = "POLITICS"
	BusinessvideosContentCategoryEnumParamScience       BusinessvideosContentCategoryEnumParam = "SCIENCE"
	BusinessvideosContentCategoryEnumParamSports        BusinessvideosContentCategoryEnumParam = "SPORTS"
	BusinessvideosContentCategoryEnumParamTechnology    BusinessvideosContentCategoryEnumParam = "TECHNOLOGY"
	BusinessvideosContentCategoryEnumParamVideoGaming   BusinessvideosContentCategoryEnumParam = "VIDEO_GAMING"
)

func (value BusinessvideosContentCategoryEnumParam) String() string {
	return string(value)
}
