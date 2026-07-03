package enums

type GroupliveVideosStatusEnumParam string

const (
	GroupliveVideosStatusEnumParamLiveNow              GroupliveVideosStatusEnumParam = "LIVE_NOW"
	GroupliveVideosStatusEnumParamScheduledCanceled    GroupliveVideosStatusEnumParam = "SCHEDULED_CANCELED"
	GroupliveVideosStatusEnumParamScheduledLive        GroupliveVideosStatusEnumParam = "SCHEDULED_LIVE"
	GroupliveVideosStatusEnumParamScheduledUnpublished GroupliveVideosStatusEnumParam = "SCHEDULED_UNPUBLISHED"
	GroupliveVideosStatusEnumParamUnpublished          GroupliveVideosStatusEnumParam = "UNPUBLISHED"
)

func (value GroupliveVideosStatusEnumParam) String() string {
	return string(value)
}
