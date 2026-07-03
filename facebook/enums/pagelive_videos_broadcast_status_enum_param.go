package enums

type PageliveVideosBroadcastStatusEnumParam string

const (
	PageliveVideosBroadcastStatusEnumParamLive                 PageliveVideosBroadcastStatusEnumParam = "LIVE"
	PageliveVideosBroadcastStatusEnumParamLiveStopped          PageliveVideosBroadcastStatusEnumParam = "LIVE_STOPPED"
	PageliveVideosBroadcastStatusEnumParamProcessing           PageliveVideosBroadcastStatusEnumParam = "PROCESSING"
	PageliveVideosBroadcastStatusEnumParamScheduledCanceled    PageliveVideosBroadcastStatusEnumParam = "SCHEDULED_CANCELED"
	PageliveVideosBroadcastStatusEnumParamScheduledExpired     PageliveVideosBroadcastStatusEnumParam = "SCHEDULED_EXPIRED"
	PageliveVideosBroadcastStatusEnumParamScheduledLive        PageliveVideosBroadcastStatusEnumParam = "SCHEDULED_LIVE"
	PageliveVideosBroadcastStatusEnumParamScheduledUnpublished PageliveVideosBroadcastStatusEnumParam = "SCHEDULED_UNPUBLISHED"
	PageliveVideosBroadcastStatusEnumParamUnpublished          PageliveVideosBroadcastStatusEnumParam = "UNPUBLISHED"
	PageliveVideosBroadcastStatusEnumParamVod                  PageliveVideosBroadcastStatusEnumParam = "VOD"
)

func (value PageliveVideosBroadcastStatusEnumParam) String() string {
	return string(value)
}
