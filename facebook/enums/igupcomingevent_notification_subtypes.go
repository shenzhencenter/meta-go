package enums

type IgupcomingeventNotificationSubtypes string

const (
	IgupcomingeventNotificationSubtypesAfterEventX1DAY   IgupcomingeventNotificationSubtypes = "AFTER_EVENT_1DAY"
	IgupcomingeventNotificationSubtypesAfterEventX2DAY   IgupcomingeventNotificationSubtypes = "AFTER_EVENT_2DAY"
	IgupcomingeventNotificationSubtypesAfterEventX3DAY   IgupcomingeventNotificationSubtypes = "AFTER_EVENT_3DAY"
	IgupcomingeventNotificationSubtypesAfterEventX4DAY   IgupcomingeventNotificationSubtypes = "AFTER_EVENT_4DAY"
	IgupcomingeventNotificationSubtypesAfterEventX5DAY   IgupcomingeventNotificationSubtypes = "AFTER_EVENT_5DAY"
	IgupcomingeventNotificationSubtypesAfterEventX6DAY   IgupcomingeventNotificationSubtypes = "AFTER_EVENT_6DAY"
	IgupcomingeventNotificationSubtypesAfterEventX7DAY   IgupcomingeventNotificationSubtypes = "AFTER_EVENT_7DAY"
	IgupcomingeventNotificationSubtypesBeforeEventX15MIN IgupcomingeventNotificationSubtypes = "BEFORE_EVENT_15MIN"
	IgupcomingeventNotificationSubtypesBeforeEventX1DAY  IgupcomingeventNotificationSubtypes = "BEFORE_EVENT_1DAY"
	IgupcomingeventNotificationSubtypesBeforeEventX1HOUR IgupcomingeventNotificationSubtypes = "BEFORE_EVENT_1HOUR"
	IgupcomingeventNotificationSubtypesBeforeEventX2DAY  IgupcomingeventNotificationSubtypes = "BEFORE_EVENT_2DAY"
	IgupcomingeventNotificationSubtypesEventStart        IgupcomingeventNotificationSubtypes = "EVENT_START"
	IgupcomingeventNotificationSubtypesRescheduled       IgupcomingeventNotificationSubtypes = "RESCHEDULED"
)

func (value IgupcomingeventNotificationSubtypes) String() string {
	return string(value)
}
