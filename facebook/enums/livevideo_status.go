package enums

type LivevideoStatus string

const (
	LivevideoStatusLiveNow              LivevideoStatus = "LIVE_NOW"
	LivevideoStatusScheduledCanceled    LivevideoStatus = "SCHEDULED_CANCELED"
	LivevideoStatusScheduledLive        LivevideoStatus = "SCHEDULED_LIVE"
	LivevideoStatusScheduledUnpublished LivevideoStatus = "SCHEDULED_UNPUBLISHED"
	LivevideoStatusUnpublished          LivevideoStatus = "UNPUBLISHED"
)

func (value LivevideoStatus) String() string {
	return string(value)
}
