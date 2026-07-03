package enums

type UserliveVideosBroadcastStatusEnumParam string

const (
	UserliveVideosBroadcastStatusEnumParamLive                 UserliveVideosBroadcastStatusEnumParam = "LIVE"
	UserliveVideosBroadcastStatusEnumParamLiveStopped          UserliveVideosBroadcastStatusEnumParam = "LIVE_STOPPED"
	UserliveVideosBroadcastStatusEnumParamProcessing           UserliveVideosBroadcastStatusEnumParam = "PROCESSING"
	UserliveVideosBroadcastStatusEnumParamScheduledCanceled    UserliveVideosBroadcastStatusEnumParam = "SCHEDULED_CANCELED"
	UserliveVideosBroadcastStatusEnumParamScheduledExpired     UserliveVideosBroadcastStatusEnumParam = "SCHEDULED_EXPIRED"
	UserliveVideosBroadcastStatusEnumParamScheduledLive        UserliveVideosBroadcastStatusEnumParam = "SCHEDULED_LIVE"
	UserliveVideosBroadcastStatusEnumParamScheduledUnpublished UserliveVideosBroadcastStatusEnumParam = "SCHEDULED_UNPUBLISHED"
	UserliveVideosBroadcastStatusEnumParamUnpublished          UserliveVideosBroadcastStatusEnumParam = "UNPUBLISHED"
	UserliveVideosBroadcastStatusEnumParamVod                  UserliveVideosBroadcastStatusEnumParam = "VOD"
)

func (value UserliveVideosBroadcastStatusEnumParam) String() string {
	return string(value)
}
