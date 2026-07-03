package enums

type UservideosUploadPhaseEnumParam string

const (
	UservideosUploadPhaseEnumParamCancel   UservideosUploadPhaseEnumParam = "cancel"
	UservideosUploadPhaseEnumParamFinish   UservideosUploadPhaseEnumParam = "finish"
	UservideosUploadPhaseEnumParamStart    UservideosUploadPhaseEnumParam = "start"
	UservideosUploadPhaseEnumParamTransfer UservideosUploadPhaseEnumParam = "transfer"
)

func (value UservideosUploadPhaseEnumParam) String() string {
	return string(value)
}
