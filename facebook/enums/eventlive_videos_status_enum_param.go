package enums

type EventliveVideosStatusEnumParam string

const (
	EventliveVideosStatusEnumParamLiveNow              EventliveVideosStatusEnumParam = "LIVE_NOW"
	EventliveVideosStatusEnumParamScheduledCanceled    EventliveVideosStatusEnumParam = "SCHEDULED_CANCELED"
	EventliveVideosStatusEnumParamScheduledLive        EventliveVideosStatusEnumParam = "SCHEDULED_LIVE"
	EventliveVideosStatusEnumParamScheduledUnpublished EventliveVideosStatusEnumParam = "SCHEDULED_UNPUBLISHED"
	EventliveVideosStatusEnumParamUnpublished          EventliveVideosStatusEnumParam = "UNPUBLISHED"
)

func (value EventliveVideosStatusEnumParam) String() string {
	return string(value)
}
