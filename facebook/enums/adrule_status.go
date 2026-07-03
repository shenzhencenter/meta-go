package enums

type AdruleStatus string

const (
	AdruleStatusDeleted   AdruleStatus = "DELETED"
	AdruleStatusDisabled  AdruleStatus = "DISABLED"
	AdruleStatusEnabled   AdruleStatus = "ENABLED"
	AdruleStatusHasIssues AdruleStatus = "HAS_ISSUES"
)

func (value AdruleStatus) String() string {
	return string(value)
}
