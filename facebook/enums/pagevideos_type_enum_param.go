package enums

type PagevideosTypeEnumParam string

const (
	PagevideosTypeEnumParamTagged   PagevideosTypeEnumParam = "TAGGED"
	PagevideosTypeEnumParamUploaded PagevideosTypeEnumParam = "UPLOADED"
)

func (value PagevideosTypeEnumParam) String() string {
	return string(value)
}
