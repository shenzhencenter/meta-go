package enums

type ProductitemStatus string

const (
	ProductitemStatusPublished ProductitemStatus = "PUBLISHED"
	ProductitemStatusStaging   ProductitemStatus = "STAGING"
)

func (value ProductitemStatus) String() string {
	return string(value)
}
