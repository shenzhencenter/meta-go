package enums

type ProductcatalogagenciesPermittedTasksEnumParam string

const (
	ProductcatalogagenciesPermittedTasksEnumParamAaAnalyze ProductcatalogagenciesPermittedTasksEnumParam = "AA_ANALYZE"
	ProductcatalogagenciesPermittedTasksEnumParamAdvertise ProductcatalogagenciesPermittedTasksEnumParam = "ADVERTISE"
	ProductcatalogagenciesPermittedTasksEnumParamManage    ProductcatalogagenciesPermittedTasksEnumParam = "MANAGE"
	ProductcatalogagenciesPermittedTasksEnumParamManageAr  ProductcatalogagenciesPermittedTasksEnumParam = "MANAGE_AR"
)

func (value ProductcatalogagenciesPermittedTasksEnumParam) String() string {
	return string(value)
}
