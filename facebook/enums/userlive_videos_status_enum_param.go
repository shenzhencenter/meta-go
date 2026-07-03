package enums

type UserliveVideosStatusEnumParam string

const (
	UserliveVideosStatusEnumParamLiveNow              UserliveVideosStatusEnumParam = "LIVE_NOW"
	UserliveVideosStatusEnumParamScheduledCanceled    UserliveVideosStatusEnumParam = "SCHEDULED_CANCELED"
	UserliveVideosStatusEnumParamScheduledLive        UserliveVideosStatusEnumParam = "SCHEDULED_LIVE"
	UserliveVideosStatusEnumParamScheduledUnpublished UserliveVideosStatusEnumParam = "SCHEDULED_UNPUBLISHED"
	UserliveVideosStatusEnumParamUnpublished          UserliveVideosStatusEnumParam = "UNPUBLISHED"
)

func (value UserliveVideosStatusEnumParam) String() string {
	return string(value)
}
