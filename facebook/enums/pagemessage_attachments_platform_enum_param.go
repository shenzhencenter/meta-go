package enums

type PagemessageAttachmentsPlatformEnumParam string

const (
	PagemessageAttachmentsPlatformEnumParamInstagram PagemessageAttachmentsPlatformEnumParam = "INSTAGRAM"
	PagemessageAttachmentsPlatformEnumParamMessenger PagemessageAttachmentsPlatformEnumParam = "MESSENGER"
)

func (value PagemessageAttachmentsPlatformEnumParam) String() string {
	return string(value)
}
