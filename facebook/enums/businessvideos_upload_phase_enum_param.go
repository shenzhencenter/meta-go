package enums

type BusinessvideosUploadPhaseEnumParam string

const (
	BusinessvideosUploadPhaseEnumParamCancel   BusinessvideosUploadPhaseEnumParam = "cancel"
	BusinessvideosUploadPhaseEnumParamFinish   BusinessvideosUploadPhaseEnumParam = "finish"
	BusinessvideosUploadPhaseEnumParamStart    BusinessvideosUploadPhaseEnumParam = "start"
	BusinessvideosUploadPhaseEnumParamTransfer BusinessvideosUploadPhaseEnumParam = "transfer"
)

func (value BusinessvideosUploadPhaseEnumParam) String() string {
	return string(value)
}
