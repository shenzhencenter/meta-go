package enums

type PagemessagesNotificationTypeEnumParam string

const (
	PagemessagesNotificationTypeEnumParamNoPush     PagemessagesNotificationTypeEnumParam = "NO_PUSH"
	PagemessagesNotificationTypeEnumParamRegular    PagemessagesNotificationTypeEnumParam = "REGULAR"
	PagemessagesNotificationTypeEnumParamSilentPush PagemessagesNotificationTypeEnumParam = "SILENT_PUSH"
)

func (value PagemessagesNotificationTypeEnumParam) String() string {
	return string(value)
}
