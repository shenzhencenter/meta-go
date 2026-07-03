package enums

type UsernotificationsTypeEnumParam string

const (
	UsernotificationsTypeEnumParamContentUpdate UsernotificationsTypeEnumParam = "content_update"
	UsernotificationsTypeEnumParamGeneric       UsernotificationsTypeEnumParam = "generic"
)

func (value UsernotificationsTypeEnumParam) String() string {
	return string(value)
}
