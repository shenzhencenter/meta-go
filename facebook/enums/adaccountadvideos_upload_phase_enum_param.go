package enums

type AdaccountadvideosUploadPhaseEnumParam string

const (
	AdaccountadvideosUploadPhaseEnumParamCancel   AdaccountadvideosUploadPhaseEnumParam = "cancel"
	AdaccountadvideosUploadPhaseEnumParamFinish   AdaccountadvideosUploadPhaseEnumParam = "finish"
	AdaccountadvideosUploadPhaseEnumParamStart    AdaccountadvideosUploadPhaseEnumParam = "start"
	AdaccountadvideosUploadPhaseEnumParamTransfer AdaccountadvideosUploadPhaseEnumParam = "transfer"
)

func (value AdaccountadvideosUploadPhaseEnumParam) String() string {
	return string(value)
}
