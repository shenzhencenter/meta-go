package enums

type UserphotosTypeEnumParam string

const (
	UserphotosTypeEnumParamTagged   UserphotosTypeEnumParam = "tagged"
	UserphotosTypeEnumParamUploaded UserphotosTypeEnumParam = "uploaded"
)

func (value UserphotosTypeEnumParam) String() string {
	return string(value)
}
