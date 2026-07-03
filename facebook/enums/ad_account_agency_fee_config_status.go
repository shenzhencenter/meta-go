package enums

type AdaccountagencyfeeconfigStatus string

const (
	AdaccountagencyfeeconfigStatusActive   AdaccountagencyfeeconfigStatus = "ACTIVE"
	AdaccountagencyfeeconfigStatusArchived AdaccountagencyfeeconfigStatus = "ARCHIVED"
	AdaccountagencyfeeconfigStatusDraft    AdaccountagencyfeeconfigStatus = "DRAFT"
	AdaccountagencyfeeconfigStatusUnknown  AdaccountagencyfeeconfigStatus = "UNKNOWN"
)

func (value AdaccountagencyfeeconfigStatus) String() string {
	return string(value)
}
