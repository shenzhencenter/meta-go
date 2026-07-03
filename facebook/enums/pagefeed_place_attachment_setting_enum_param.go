package enums

type PagefeedPlaceAttachmentSettingEnumParam string

const (
	PagefeedPlaceAttachmentSettingEnumParamX1 PagefeedPlaceAttachmentSettingEnumParam = "1"
	PagefeedPlaceAttachmentSettingEnumParamX2 PagefeedPlaceAttachmentSettingEnumParam = "2"
)

func (value PagefeedPlaceAttachmentSettingEnumParam) String() string {
	return string(value)
}
