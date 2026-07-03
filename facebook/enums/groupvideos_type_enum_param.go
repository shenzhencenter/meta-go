package enums

type GroupvideosTypeEnumParam string

const (
	GroupvideosTypeEnumParamTagged   GroupvideosTypeEnumParam = "tagged"
	GroupvideosTypeEnumParamUploaded GroupvideosTypeEnumParam = "uploaded"
)

func (value GroupvideosTypeEnumParam) String() string {
	return string(value)
}
