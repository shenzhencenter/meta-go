package enums

type AdgroupStatus string

const (
	AdgroupStatusActive   AdgroupStatus = "ACTIVE"
	AdgroupStatusArchived AdgroupStatus = "ARCHIVED"
	AdgroupStatusDeleted  AdgroupStatus = "DELETED"
	AdgroupStatusPaused   AdgroupStatus = "PAUSED"
)

func (value AdgroupStatus) String() string {
	return string(value)
}
