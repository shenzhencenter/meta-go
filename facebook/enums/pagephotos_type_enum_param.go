package enums

type PagephotosTypeEnumParam string

const (
	PagephotosTypeEnumParamProfile  PagephotosTypeEnumParam = "profile"
	PagephotosTypeEnumParamTagged   PagephotosTypeEnumParam = "tagged"
	PagephotosTypeEnumParamUploaded PagephotosTypeEnumParam = "uploaded"
)

func (value PagephotosTypeEnumParam) String() string {
	return string(value)
}
