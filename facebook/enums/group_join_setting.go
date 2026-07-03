package enums

type GroupJoinSetting string

const (
	GroupJoinSettingAdminOnly GroupJoinSetting = "ADMIN_ONLY"
	GroupJoinSettingAnyone    GroupJoinSetting = "ANYONE"
	GroupJoinSettingNone      GroupJoinSetting = "NONE"
)

func (value GroupJoinSetting) String() string {
	return string(value)
}
