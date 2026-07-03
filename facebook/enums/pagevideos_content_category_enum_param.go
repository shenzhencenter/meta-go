package enums

type PagevideosContentCategoryEnumParam string

const (
	PagevideosContentCategoryEnumParamBeautyFashion PagevideosContentCategoryEnumParam = "BEAUTY_FASHION"
	PagevideosContentCategoryEnumParamBusiness      PagevideosContentCategoryEnumParam = "BUSINESS"
	PagevideosContentCategoryEnumParamCarsTrucks    PagevideosContentCategoryEnumParam = "CARS_TRUCKS"
	PagevideosContentCategoryEnumParamComedy        PagevideosContentCategoryEnumParam = "COMEDY"
	PagevideosContentCategoryEnumParamCuteAnimals   PagevideosContentCategoryEnumParam = "CUTE_ANIMALS"
	PagevideosContentCategoryEnumParamEntertainment PagevideosContentCategoryEnumParam = "ENTERTAINMENT"
	PagevideosContentCategoryEnumParamFamily        PagevideosContentCategoryEnumParam = "FAMILY"
	PagevideosContentCategoryEnumParamFoodHealth    PagevideosContentCategoryEnumParam = "FOOD_HEALTH"
	PagevideosContentCategoryEnumParamHome          PagevideosContentCategoryEnumParam = "HOME"
	PagevideosContentCategoryEnumParamLifestyle     PagevideosContentCategoryEnumParam = "LIFESTYLE"
	PagevideosContentCategoryEnumParamMusic         PagevideosContentCategoryEnumParam = "MUSIC"
	PagevideosContentCategoryEnumParamNews          PagevideosContentCategoryEnumParam = "NEWS"
	PagevideosContentCategoryEnumParamOther         PagevideosContentCategoryEnumParam = "OTHER"
	PagevideosContentCategoryEnumParamPolitics      PagevideosContentCategoryEnumParam = "POLITICS"
	PagevideosContentCategoryEnumParamScience       PagevideosContentCategoryEnumParam = "SCIENCE"
	PagevideosContentCategoryEnumParamSports        PagevideosContentCategoryEnumParam = "SPORTS"
	PagevideosContentCategoryEnumParamTechnology    PagevideosContentCategoryEnumParam = "TECHNOLOGY"
	PagevideosContentCategoryEnumParamVideoGaming   PagevideosContentCategoryEnumParam = "VIDEO_GAMING"
)

func (value PagevideosContentCategoryEnumParam) String() string {
	return string(value)
}
