package enums

type ProductcatalogeventStatsBreakdownsEnumParam string

const (
	ProductcatalogeventStatsBreakdownsEnumParamDeviceType ProductcatalogeventStatsBreakdownsEnumParam = "DEVICE_TYPE"
)

func (value ProductcatalogeventStatsBreakdownsEnumParam) String() string {
	return string(value)
}
