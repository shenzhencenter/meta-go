package enums

type PagevideoReelsUploadPhaseEnumParam string

const (
	PagevideoReelsUploadPhaseEnumParamFinish PagevideoReelsUploadPhaseEnumParam = "FINISH"
	PagevideoReelsUploadPhaseEnumParamStart  PagevideoReelsUploadPhaseEnumParam = "START"
)

func (value PagevideoReelsUploadPhaseEnumParam) String() string {
	return string(value)
}
