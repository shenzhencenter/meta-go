package enums

type AdStatus string

const (
	AdStatusActive   AdStatus = "ACTIVE"
	AdStatusArchived AdStatus = "ARCHIVED"
	AdStatusDeleted  AdStatus = "DELETED"
	AdStatusPaused   AdStatus = "PAUSED"
)

func (value AdStatus) String() string {
	return string(value)
}
