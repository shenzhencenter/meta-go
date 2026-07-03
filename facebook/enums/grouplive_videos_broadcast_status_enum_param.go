package enums

type GroupliveVideosBroadcastStatusEnumParam string

const (
	GroupliveVideosBroadcastStatusEnumParamLive                 GroupliveVideosBroadcastStatusEnumParam = "LIVE"
	GroupliveVideosBroadcastStatusEnumParamLiveStopped          GroupliveVideosBroadcastStatusEnumParam = "LIVE_STOPPED"
	GroupliveVideosBroadcastStatusEnumParamProcessing           GroupliveVideosBroadcastStatusEnumParam = "PROCESSING"
	GroupliveVideosBroadcastStatusEnumParamScheduledCanceled    GroupliveVideosBroadcastStatusEnumParam = "SCHEDULED_CANCELED"
	GroupliveVideosBroadcastStatusEnumParamScheduledExpired     GroupliveVideosBroadcastStatusEnumParam = "SCHEDULED_EXPIRED"
	GroupliveVideosBroadcastStatusEnumParamScheduledLive        GroupliveVideosBroadcastStatusEnumParam = "SCHEDULED_LIVE"
	GroupliveVideosBroadcastStatusEnumParamScheduledUnpublished GroupliveVideosBroadcastStatusEnumParam = "SCHEDULED_UNPUBLISHED"
	GroupliveVideosBroadcastStatusEnumParamUnpublished          GroupliveVideosBroadcastStatusEnumParam = "UNPUBLISHED"
	GroupliveVideosBroadcastStatusEnumParamVod                  GroupliveVideosBroadcastStatusEnumParam = "VOD"
)

func (value GroupliveVideosBroadcastStatusEnumParam) String() string {
	return string(value)
}
