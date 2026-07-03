package enums

type AdasyncrequestsetNotificationModeEnum string

const (
	AdasyncrequestsetNotificationModeEnumOff        AdasyncrequestsetNotificationModeEnum = "OFF"
	AdasyncrequestsetNotificationModeEnumOnComplete AdasyncrequestsetNotificationModeEnum = "ON_COMPLETE"
)

func (value AdasyncrequestsetNotificationModeEnum) String() string {
	return string(value)
}
