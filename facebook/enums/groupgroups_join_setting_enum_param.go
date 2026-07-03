package enums

type GroupgroupsJoinSettingEnumParam string

const (
	GroupgroupsJoinSettingEnumParamAdminOnly GroupgroupsJoinSettingEnumParam = "ADMIN_ONLY"
	GroupgroupsJoinSettingEnumParamAnyone    GroupgroupsJoinSettingEnumParam = "ANYONE"
	GroupgroupsJoinSettingEnumParamNone      GroupgroupsJoinSettingEnumParam = "NONE"
)

func (value GroupgroupsJoinSettingEnumParam) String() string {
	return string(value)
}
