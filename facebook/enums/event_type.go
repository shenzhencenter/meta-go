package enums

type EventType string

const (
	EventTypeCommunity          EventType = "community"
	EventTypeFriends            EventType = "friends"
	EventTypeGroup              EventType = "group"
	EventTypeMessengerCommunity EventType = "messenger_community"
	EventTypePrivate            EventType = "private"
	EventTypePublic             EventType = "public"
	EventTypeWorkCompany        EventType = "work_company"
)

func (value EventType) String() string {
	return string(value)
}
