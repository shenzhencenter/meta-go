package enums

type PagevideosUploadPhaseEnumParam string

const (
	PagevideosUploadPhaseEnumParamCancel   PagevideosUploadPhaseEnumParam = "cancel"
	PagevideosUploadPhaseEnumParamFinish   PagevideosUploadPhaseEnumParam = "finish"
	PagevideosUploadPhaseEnumParamStart    PagevideosUploadPhaseEnumParam = "start"
	PagevideosUploadPhaseEnumParamTransfer PagevideosUploadPhaseEnumParam = "transfer"
)

func (value PagevideosUploadPhaseEnumParam) String() string {
	return string(value)
}
