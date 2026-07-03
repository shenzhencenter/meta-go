package enums

type UserpermissionsStatusEnumParam string

const (
	UserpermissionsStatusEnumParamDeclined UserpermissionsStatusEnumParam = "declined"
	UserpermissionsStatusEnumParamExpired  UserpermissionsStatusEnumParam = "expired"
	UserpermissionsStatusEnumParamGranted  UserpermissionsStatusEnumParam = "granted"
)

func (value UserpermissionsStatusEnumParam) String() string {
	return string(value)
}
