package enums

type OfflineproductitemVisibility string

const (
	OfflineproductitemVisibilityPublished OfflineproductitemVisibility = "PUBLISHED"
	OfflineproductitemVisibilityStaging   OfflineproductitemVisibility = "STAGING"
)

func (value OfflineproductitemVisibility) String() string {
	return string(value)
}
