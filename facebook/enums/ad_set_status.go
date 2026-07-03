package enums

type AdsetStatus string

const (
	AdsetStatusActive   AdsetStatus = "ACTIVE"
	AdsetStatusArchived AdsetStatus = "ARCHIVED"
	AdsetStatusDeleted  AdsetStatus = "DELETED"
	AdsetStatusPaused   AdsetStatus = "PAUSED"
)

func (value AdsetStatus) String() string {
	return string(value)
}
