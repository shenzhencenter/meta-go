package enums

type PlacetopicIconSize string

const (
	PlacetopicIconSizeX24 PlacetopicIconSize = "24"
	PlacetopicIconSizeX36 PlacetopicIconSize = "36"
	PlacetopicIconSizeX48 PlacetopicIconSize = "48"
	PlacetopicIconSizeX72 PlacetopicIconSize = "72"
)

func (value PlacetopicIconSize) String() string {
	return string(value)
}
