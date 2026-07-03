package enums

type AdspixelofflineEventUploadsOrderEnumParam string

const (
	AdspixelofflineEventUploadsOrderEnumParamAscending  AdspixelofflineEventUploadsOrderEnumParam = "ASCENDING"
	AdspixelofflineEventUploadsOrderEnumParamDescending AdspixelofflineEventUploadsOrderEnumParam = "DESCENDING"
)

func (value AdspixelofflineEventUploadsOrderEnumParam) String() string {
	return string(value)
}
