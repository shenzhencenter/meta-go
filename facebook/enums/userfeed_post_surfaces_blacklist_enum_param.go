package enums

type UserfeedPostSurfacesBlacklistEnumParam string

const (
	UserfeedPostSurfacesBlacklistEnumParamX1 UserfeedPostSurfacesBlacklistEnumParam = "1"
	UserfeedPostSurfacesBlacklistEnumParamX2 UserfeedPostSurfacesBlacklistEnumParam = "2"
	UserfeedPostSurfacesBlacklistEnumParamX3 UserfeedPostSurfacesBlacklistEnumParam = "3"
	UserfeedPostSurfacesBlacklistEnumParamX4 UserfeedPostSurfacesBlacklistEnumParam = "4"
	UserfeedPostSurfacesBlacklistEnumParamX5 UserfeedPostSurfacesBlacklistEnumParam = "5"
)

func (value UserfeedPostSurfacesBlacklistEnumParam) String() string {
	return string(value)
}
