package enums

type GroupPostPermissions string

const (
	GroupPostPermissionsAdminOnly GroupPostPermissions = "ADMIN_ONLY"
	GroupPostPermissionsAnyone    GroupPostPermissions = "ANYONE"
	GroupPostPermissionsNone      GroupPostPermissions = "NONE"
)

func (value GroupPostPermissions) String() string {
	return string(value)
}
