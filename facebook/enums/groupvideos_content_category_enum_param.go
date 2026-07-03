package enums

type GroupvideosContentCategoryEnumParam string

const (
	GroupvideosContentCategoryEnumParamBeautyFashion GroupvideosContentCategoryEnumParam = "BEAUTY_FASHION"
	GroupvideosContentCategoryEnumParamBusiness      GroupvideosContentCategoryEnumParam = "BUSINESS"
	GroupvideosContentCategoryEnumParamCarsTrucks    GroupvideosContentCategoryEnumParam = "CARS_TRUCKS"
	GroupvideosContentCategoryEnumParamComedy        GroupvideosContentCategoryEnumParam = "COMEDY"
	GroupvideosContentCategoryEnumParamCuteAnimals   GroupvideosContentCategoryEnumParam = "CUTE_ANIMALS"
	GroupvideosContentCategoryEnumParamEntertainment GroupvideosContentCategoryEnumParam = "ENTERTAINMENT"
	GroupvideosContentCategoryEnumParamFamily        GroupvideosContentCategoryEnumParam = "FAMILY"
	GroupvideosContentCategoryEnumParamFoodHealth    GroupvideosContentCategoryEnumParam = "FOOD_HEALTH"
	GroupvideosContentCategoryEnumParamHome          GroupvideosContentCategoryEnumParam = "HOME"
	GroupvideosContentCategoryEnumParamLifestyle     GroupvideosContentCategoryEnumParam = "LIFESTYLE"
	GroupvideosContentCategoryEnumParamMusic         GroupvideosContentCategoryEnumParam = "MUSIC"
	GroupvideosContentCategoryEnumParamNews          GroupvideosContentCategoryEnumParam = "NEWS"
	GroupvideosContentCategoryEnumParamOther         GroupvideosContentCategoryEnumParam = "OTHER"
	GroupvideosContentCategoryEnumParamPolitics      GroupvideosContentCategoryEnumParam = "POLITICS"
	GroupvideosContentCategoryEnumParamScience       GroupvideosContentCategoryEnumParam = "SCIENCE"
	GroupvideosContentCategoryEnumParamSports        GroupvideosContentCategoryEnumParam = "SPORTS"
	GroupvideosContentCategoryEnumParamTechnology    GroupvideosContentCategoryEnumParam = "TECHNOLOGY"
	GroupvideosContentCategoryEnumParamVideoGaming   GroupvideosContentCategoryEnumParam = "VIDEO_GAMING"
)

func (value GroupvideosContentCategoryEnumParam) String() string {
	return string(value)
}
