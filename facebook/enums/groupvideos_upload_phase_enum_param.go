package enums

type GroupvideosUploadPhaseEnumParam string

const (
	GroupvideosUploadPhaseEnumParamCancel   GroupvideosUploadPhaseEnumParam = "cancel"
	GroupvideosUploadPhaseEnumParamFinish   GroupvideosUploadPhaseEnumParam = "finish"
	GroupvideosUploadPhaseEnumParamStart    GroupvideosUploadPhaseEnumParam = "start"
	GroupvideosUploadPhaseEnumParamTransfer GroupvideosUploadPhaseEnumParam = "transfer"
)

func (value GroupvideosUploadPhaseEnumParam) String() string {
	return string(value)
}
