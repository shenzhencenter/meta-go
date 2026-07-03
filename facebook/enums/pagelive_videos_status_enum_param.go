package enums

type PageliveVideosStatusEnumParam string

const (
	PageliveVideosStatusEnumParamLiveNow              PageliveVideosStatusEnumParam = "LIVE_NOW"
	PageliveVideosStatusEnumParamScheduledCanceled    PageliveVideosStatusEnumParam = "SCHEDULED_CANCELED"
	PageliveVideosStatusEnumParamScheduledLive        PageliveVideosStatusEnumParam = "SCHEDULED_LIVE"
	PageliveVideosStatusEnumParamScheduledUnpublished PageliveVideosStatusEnumParam = "SCHEDULED_UNPUBLISHED"
	PageliveVideosStatusEnumParamUnpublished          PageliveVideosStatusEnumParam = "UNPUBLISHED"
)

func (value PageliveVideosStatusEnumParam) String() string {
	return string(value)
}
