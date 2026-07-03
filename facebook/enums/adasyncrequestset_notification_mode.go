package enums

type AdasyncrequestsetNotificationMode string

const (
	AdasyncrequestsetNotificationModeOff        AdasyncrequestsetNotificationMode = "OFF"
	AdasyncrequestsetNotificationModeOnComplete AdasyncrequestsetNotificationMode = "ON_COMPLETE"
)

func (value AdasyncrequestsetNotificationMode) String() string {
	return string(value)
}
