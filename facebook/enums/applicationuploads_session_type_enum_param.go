package enums

type ApplicationuploadsSessionTypeEnumParam string

const (
	ApplicationuploadsSessionTypeEnumParamAttachment ApplicationuploadsSessionTypeEnumParam = "attachment"
)

func (value ApplicationuploadsSessionTypeEnumParam) String() string {
	return string(value)
}
