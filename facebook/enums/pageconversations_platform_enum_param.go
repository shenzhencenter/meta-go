package enums

type PageconversationsPlatformEnumParam string

const (
	PageconversationsPlatformEnumParamInstagram PageconversationsPlatformEnumParam = "INSTAGRAM"
	PageconversationsPlatformEnumParamMessenger PageconversationsPlatformEnumParam = "MESSENGER"
)

func (value PageconversationsPlatformEnumParam) String() string {
	return string(value)
}
