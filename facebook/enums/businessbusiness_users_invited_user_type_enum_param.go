package enums

type BusinessbusinessUsersInvitedUserTypeEnumParam string

const (
	BusinessbusinessUsersInvitedUserTypeEnumParamFb  BusinessbusinessUsersInvitedUserTypeEnumParam = "FB"
	BusinessbusinessUsersInvitedUserTypeEnumParamMwa BusinessbusinessUsersInvitedUserTypeEnumParam = "MWA"
)

func (value BusinessbusinessUsersInvitedUserTypeEnumParam) String() string {
	return string(value)
}
