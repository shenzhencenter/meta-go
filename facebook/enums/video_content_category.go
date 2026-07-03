package enums

type VideoContentCategory string

const (
	VideoContentCategoryBeautyFashion VideoContentCategory = "BEAUTY_FASHION"
	VideoContentCategoryBusiness      VideoContentCategory = "BUSINESS"
	VideoContentCategoryCarsTrucks    VideoContentCategory = "CARS_TRUCKS"
	VideoContentCategoryComedy        VideoContentCategory = "COMEDY"
	VideoContentCategoryCuteAnimals   VideoContentCategory = "CUTE_ANIMALS"
	VideoContentCategoryEntertainment VideoContentCategory = "ENTERTAINMENT"
	VideoContentCategoryFamily        VideoContentCategory = "FAMILY"
	VideoContentCategoryFoodHealth    VideoContentCategory = "FOOD_HEALTH"
	VideoContentCategoryHome          VideoContentCategory = "HOME"
	VideoContentCategoryLifestyle     VideoContentCategory = "LIFESTYLE"
	VideoContentCategoryMusic         VideoContentCategory = "MUSIC"
	VideoContentCategoryNews          VideoContentCategory = "NEWS"
	VideoContentCategoryOther         VideoContentCategory = "OTHER"
	VideoContentCategoryPolitics      VideoContentCategory = "POLITICS"
	VideoContentCategoryScience       VideoContentCategory = "SCIENCE"
	VideoContentCategorySports        VideoContentCategory = "SPORTS"
	VideoContentCategoryTechnology    VideoContentCategory = "TECHNOLOGY"
	VideoContentCategoryVideoGaming   VideoContentCategory = "VIDEO_GAMING"
)

func (value VideoContentCategory) String() string {
	return string(value)
}
