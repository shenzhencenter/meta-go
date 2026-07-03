package enums

type PagevideoStoriesUploadPhaseEnumParam string

const (
	PagevideoStoriesUploadPhaseEnumParamFinish PagevideoStoriesUploadPhaseEnumParam = "FINISH"
	PagevideoStoriesUploadPhaseEnumParamStart  PagevideoStoriesUploadPhaseEnumParam = "START"
)

func (value PagevideoStoriesUploadPhaseEnumParam) String() string {
	return string(value)
}
