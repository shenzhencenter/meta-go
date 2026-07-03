package enums

type AdsetConfiguredStatus string

const (
	AdsetConfiguredStatusActive   AdsetConfiguredStatus = "ACTIVE"
	AdsetConfiguredStatusArchived AdsetConfiguredStatus = "ARCHIVED"
	AdsetConfiguredStatusDeleted  AdsetConfiguredStatus = "DELETED"
	AdsetConfiguredStatusPaused   AdsetConfiguredStatus = "PAUSED"
)

func (value AdsetConfiguredStatus) String() string {
	return string(value)
}
