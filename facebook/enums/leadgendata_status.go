package enums

type LeadgendataStatus string

const (
	LeadgendataStatusActive   LeadgendataStatus = "ACTIVE"
	LeadgendataStatusArchived LeadgendataStatus = "ARCHIVED"
	LeadgendataStatusDeleted  LeadgendataStatus = "DELETED"
	LeadgendataStatusDraft    LeadgendataStatus = "DRAFT"
)

func (value LeadgendataStatus) String() string {
	return string(value)
}
