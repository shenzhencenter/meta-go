package enums

type AdConfiguredStatus string

const (
	AdConfiguredStatusActive   AdConfiguredStatus = "ACTIVE"
	AdConfiguredStatusArchived AdConfiguredStatus = "ARCHIVED"
	AdConfiguredStatusDeleted  AdConfiguredStatus = "DELETED"
	AdConfiguredStatusPaused   AdConfiguredStatus = "PAUSED"
)

func (value AdConfiguredStatus) String() string {
	return string(value)
}
