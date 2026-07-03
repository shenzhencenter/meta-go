package enums

type EventOnlineEventFormat string

const (
	EventOnlineEventFormatFbLive        EventOnlineEventFormat = "fb_live"
	EventOnlineEventFormatHorizonEvent  EventOnlineEventFormat = "horizon_event"
	EventOnlineEventFormatMessengerRoom EventOnlineEventFormat = "messenger_room"
	EventOnlineEventFormatNone          EventOnlineEventFormat = "none"
	EventOnlineEventFormatOther         EventOnlineEventFormat = "other"
	EventOnlineEventFormatThirdParty    EventOnlineEventFormat = "third_party"
)

func (value EventOnlineEventFormat) String() string {
	return string(value)
}
