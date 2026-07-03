package enums

type AdcreativeStatus string

const (
	AdcreativeStatusActive     AdcreativeStatus = "ACTIVE"
	AdcreativeStatusDeleted    AdcreativeStatus = "DELETED"
	AdcreativeStatusInProcess  AdcreativeStatus = "IN_PROCESS"
	AdcreativeStatusWithIssues AdcreativeStatus = "WITH_ISSUES"
)

func (value AdcreativeStatus) String() string {
	return string(value)
}
