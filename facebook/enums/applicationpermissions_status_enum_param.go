package enums

type ApplicationpermissionsStatusEnumParam string

const (
	ApplicationpermissionsStatusEnumParamLive       ApplicationpermissionsStatusEnumParam = "live"
	ApplicationpermissionsStatusEnumParamUnapproved ApplicationpermissionsStatusEnumParam = "unapproved"
)

func (value ApplicationpermissionsStatusEnumParam) String() string {
	return string(value)
}
