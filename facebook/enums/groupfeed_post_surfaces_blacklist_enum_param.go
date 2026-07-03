package enums

type GroupfeedPostSurfacesBlacklistEnumParam string

const (
	GroupfeedPostSurfacesBlacklistEnumParamX1 GroupfeedPostSurfacesBlacklistEnumParam = "1"
	GroupfeedPostSurfacesBlacklistEnumParamX2 GroupfeedPostSurfacesBlacklistEnumParam = "2"
	GroupfeedPostSurfacesBlacklistEnumParamX3 GroupfeedPostSurfacesBlacklistEnumParam = "3"
	GroupfeedPostSurfacesBlacklistEnumParamX4 GroupfeedPostSurfacesBlacklistEnumParam = "4"
	GroupfeedPostSurfacesBlacklistEnumParamX5 GroupfeedPostSurfacesBlacklistEnumParam = "5"
)

func (value GroupfeedPostSurfacesBlacklistEnumParam) String() string {
	return string(value)
}
