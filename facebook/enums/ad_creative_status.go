package enums

type AdcreativeStatusEnum string

const (
	AdcreativeStatusEnumActive     AdcreativeStatusEnum = "ACTIVE"
	AdcreativeStatusEnumDeleted    AdcreativeStatusEnum = "DELETED"
	AdcreativeStatusEnumInProcess  AdcreativeStatusEnum = "IN_PROCESS"
	AdcreativeStatusEnumWithIssues AdcreativeStatusEnum = "WITH_ISSUES"
)

func (value AdcreativeStatusEnum) String() string {
	return string(value)
}
