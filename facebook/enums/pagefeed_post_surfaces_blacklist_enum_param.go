package enums

type PagefeedPostSurfacesBlacklistEnumParam string

const (
	PagefeedPostSurfacesBlacklistEnumParamX1 PagefeedPostSurfacesBlacklistEnumParam = "1"
	PagefeedPostSurfacesBlacklistEnumParamX2 PagefeedPostSurfacesBlacklistEnumParam = "2"
	PagefeedPostSurfacesBlacklistEnumParamX3 PagefeedPostSurfacesBlacklistEnumParam = "3"
	PagefeedPostSurfacesBlacklistEnumParamX4 PagefeedPostSurfacesBlacklistEnumParam = "4"
	PagefeedPostSurfacesBlacklistEnumParamX5 PagefeedPostSurfacesBlacklistEnumParam = "5"
)

func (value PagefeedPostSurfacesBlacklistEnumParam) String() string {
	return string(value)
}
