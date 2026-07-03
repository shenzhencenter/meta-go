package enums

type UsernotificationsFilteringEnumParam string

const (
	UsernotificationsFilteringEnumParamEma          UsernotificationsFilteringEnumParam = "ema"
	UsernotificationsFilteringEnumParamGroups       UsernotificationsFilteringEnumParam = "groups"
	UsernotificationsFilteringEnumParamGroupsSocial UsernotificationsFilteringEnumParam = "groups_social"
)

func (value UsernotificationsFilteringEnumParam) String() string {
	return string(value)
}
