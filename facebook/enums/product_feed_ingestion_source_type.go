package enums

type ProductfeedIngestionSourceType string

const (
	ProductfeedIngestionSourceTypePrimaryFeed       ProductfeedIngestionSourceType = "primary_feed"
	ProductfeedIngestionSourceTypeSupplementaryFeed ProductfeedIngestionSourceType = "supplementary_feed"
)

func (value ProductfeedIngestionSourceType) String() string {
	return string(value)
}
