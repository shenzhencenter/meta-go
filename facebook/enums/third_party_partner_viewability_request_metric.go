package enums

type ThirdpartypartnerviewabilityrequestMetric string

const (
	ThirdpartypartnerviewabilityrequestMetricDisplayEvent ThirdpartypartnerviewabilityrequestMetric = "DISPLAY_EVENT"
	ThirdpartypartnerviewabilityrequestMetricImpression   ThirdpartypartnerviewabilityrequestMetric = "IMPRESSION"
	ThirdpartypartnerviewabilityrequestMetricVideoEvent   ThirdpartypartnerviewabilityrequestMetric = "VIDEO_EVENT"
)

func (value ThirdpartypartnerviewabilityrequestMetric) String() string {
	return string(value)
}
