package enums

type UserconversationsPlatformEnumParam string

const (
	UserconversationsPlatformEnumParamInstagram UserconversationsPlatformEnumParam = "INSTAGRAM"
	UserconversationsPlatformEnumParamMessenger UserconversationsPlatformEnumParam = "MESSENGER"
)

func (value UserconversationsPlatformEnumParam) String() string {
	return string(value)
}
