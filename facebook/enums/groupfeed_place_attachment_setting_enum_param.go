package enums

type GroupfeedPlaceAttachmentSettingEnumParam string

const (
	GroupfeedPlaceAttachmentSettingEnumParamX1 GroupfeedPlaceAttachmentSettingEnumParam = "1"
	GroupfeedPlaceAttachmentSettingEnumParamX2 GroupfeedPlaceAttachmentSettingEnumParam = "2"
)

func (value GroupfeedPlaceAttachmentSettingEnumParam) String() string {
	return string(value)
}
