package enums

type PagevideoCopyrightsContentCategoryEnumParam string

const (
	PagevideoCopyrightsContentCategoryEnumParamEpisode PagevideoCopyrightsContentCategoryEnumParam = "episode"
	PagevideoCopyrightsContentCategoryEnumParamMovie   PagevideoCopyrightsContentCategoryEnumParam = "movie"
	PagevideoCopyrightsContentCategoryEnumParamWeb     PagevideoCopyrightsContentCategoryEnumParam = "web"
)

func (value PagevideoCopyrightsContentCategoryEnumParam) String() string {
	return string(value)
}
