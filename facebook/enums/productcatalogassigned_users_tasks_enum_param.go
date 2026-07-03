package enums

type ProductcatalogassignedUsersTasksEnumParam string

const (
	ProductcatalogassignedUsersTasksEnumParamAaAnalyze ProductcatalogassignedUsersTasksEnumParam = "AA_ANALYZE"
	ProductcatalogassignedUsersTasksEnumParamAdvertise ProductcatalogassignedUsersTasksEnumParam = "ADVERTISE"
	ProductcatalogassignedUsersTasksEnumParamManage    ProductcatalogassignedUsersTasksEnumParam = "MANAGE"
	ProductcatalogassignedUsersTasksEnumParamManageAr  ProductcatalogassignedUsersTasksEnumParam = "MANAGE_AR"
)

func (value ProductcatalogassignedUsersTasksEnumParam) String() string {
	return string(value)
}
