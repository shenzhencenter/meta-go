package enums

type AdaccountadrulesLibraryStatusEnumParam string

const (
	AdaccountadrulesLibraryStatusEnumParamDeleted   AdaccountadrulesLibraryStatusEnumParam = "DELETED"
	AdaccountadrulesLibraryStatusEnumParamDisabled  AdaccountadrulesLibraryStatusEnumParam = "DISABLED"
	AdaccountadrulesLibraryStatusEnumParamEnabled   AdaccountadrulesLibraryStatusEnumParam = "ENABLED"
	AdaccountadrulesLibraryStatusEnumParamHasIssues AdaccountadrulesLibraryStatusEnumParam = "HAS_ISSUES"
)

func (value AdaccountadrulesLibraryStatusEnumParam) String() string {
	return string(value)
}
