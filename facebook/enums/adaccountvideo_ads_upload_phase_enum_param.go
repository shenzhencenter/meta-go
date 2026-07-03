package enums

type AdaccountvideoAdsUploadPhaseEnumParam string

const (
	AdaccountvideoAdsUploadPhaseEnumParamFinish AdaccountvideoAdsUploadPhaseEnumParam = "FINISH"
	AdaccountvideoAdsUploadPhaseEnumParamStart  AdaccountvideoAdsUploadPhaseEnumParam = "START"
)

func (value AdaccountvideoAdsUploadPhaseEnumParam) String() string {
	return string(value)
}
