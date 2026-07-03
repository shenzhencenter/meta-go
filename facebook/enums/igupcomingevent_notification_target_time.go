package enums

type IgupcomingeventNotificationTargetTime string

const (
	IgupcomingeventNotificationTargetTimeEventEnd   IgupcomingeventNotificationTargetTime = "EVENT_END"
	IgupcomingeventNotificationTargetTimeEventStart IgupcomingeventNotificationTargetTime = "EVENT_START"
)

func (value IgupcomingeventNotificationTargetTime) String() string {
	return string(value)
}
