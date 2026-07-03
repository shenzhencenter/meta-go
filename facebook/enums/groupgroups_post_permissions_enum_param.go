package enums

type GroupgroupsPostPermissionsEnumParam string

const (
	GroupgroupsPostPermissionsEnumParamAdminOnly GroupgroupsPostPermissionsEnumParam = "ADMIN_ONLY"
	GroupgroupsPostPermissionsEnumParamAnyone    GroupgroupsPostPermissionsEnumParam = "ANYONE"
	GroupgroupsPostPermissionsEnumParamNone      GroupgroupsPostPermissionsEnumParam = "NONE"
)

func (value GroupgroupsPostPermissionsEnumParam) String() string {
	return string(value)
}
