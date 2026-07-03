package enums

type UserfeedPlaceAttachmentSettingEnumParam string

const (
	UserfeedPlaceAttachmentSettingEnumParamX1 UserfeedPlaceAttachmentSettingEnumParam = "1"
	UserfeedPlaceAttachmentSettingEnumParamX2 UserfeedPlaceAttachmentSettingEnumParam = "2"
)

func (value UserfeedPlaceAttachmentSettingEnumParam) String() string {
	return string(value)
}
