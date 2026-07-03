package enums

type BusinessadsDatasetSortByEnumParam string

const (
	BusinessadsDatasetSortByEnumParamLastFiredTime BusinessadsDatasetSortByEnumParam = "LAST_FIRED_TIME"
	BusinessadsDatasetSortByEnumParamName          BusinessadsDatasetSortByEnumParam = "NAME"
)

func (value BusinessadsDatasetSortByEnumParam) String() string {
	return string(value)
}
