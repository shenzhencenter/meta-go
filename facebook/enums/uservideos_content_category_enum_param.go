package enums

type UservideosContentCategoryEnumParam string

const (
	UservideosContentCategoryEnumParamBeautyFashion UservideosContentCategoryEnumParam = "BEAUTY_FASHION"
	UservideosContentCategoryEnumParamBusiness      UservideosContentCategoryEnumParam = "BUSINESS"
	UservideosContentCategoryEnumParamCarsTrucks    UservideosContentCategoryEnumParam = "CARS_TRUCKS"
	UservideosContentCategoryEnumParamComedy        UservideosContentCategoryEnumParam = "COMEDY"
	UservideosContentCategoryEnumParamCuteAnimals   UservideosContentCategoryEnumParam = "CUTE_ANIMALS"
	UservideosContentCategoryEnumParamEntertainment UservideosContentCategoryEnumParam = "ENTERTAINMENT"
	UservideosContentCategoryEnumParamFamily        UservideosContentCategoryEnumParam = "FAMILY"
	UservideosContentCategoryEnumParamFoodHealth    UservideosContentCategoryEnumParam = "FOOD_HEALTH"
	UservideosContentCategoryEnumParamHome          UservideosContentCategoryEnumParam = "HOME"
	UservideosContentCategoryEnumParamLifestyle     UservideosContentCategoryEnumParam = "LIFESTYLE"
	UservideosContentCategoryEnumParamMusic         UservideosContentCategoryEnumParam = "MUSIC"
	UservideosContentCategoryEnumParamNews          UservideosContentCategoryEnumParam = "NEWS"
	UservideosContentCategoryEnumParamOther         UservideosContentCategoryEnumParam = "OTHER"
	UservideosContentCategoryEnumParamPolitics      UservideosContentCategoryEnumParam = "POLITICS"
	UservideosContentCategoryEnumParamScience       UservideosContentCategoryEnumParam = "SCIENCE"
	UservideosContentCategoryEnumParamSports        UservideosContentCategoryEnumParam = "SPORTS"
	UservideosContentCategoryEnumParamTechnology    UservideosContentCategoryEnumParam = "TECHNOLOGY"
	UservideosContentCategoryEnumParamVideoGaming   UservideosContentCategoryEnumParam = "VIDEO_GAMING"
)

func (value UservideosContentCategoryEnumParam) String() string {
	return string(value)
}
