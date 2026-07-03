package enums

type UservideosTypeEnumParam string

const (
	UservideosTypeEnumParamTagged   UservideosTypeEnumParam = "TAGGED"
	UservideosTypeEnumParamUploaded UservideosTypeEnumParam = "UPLOADED"
)

func (value UservideosTypeEnumParam) String() string {
	return string(value)
}
