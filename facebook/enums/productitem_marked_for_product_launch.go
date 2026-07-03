package enums

type ProductitemMarkedForProductLaunch string

const (
	ProductitemMarkedForProductLaunchDefault   ProductitemMarkedForProductLaunch = "default"
	ProductitemMarkedForProductLaunchMarked    ProductitemMarkedForProductLaunch = "marked"
	ProductitemMarkedForProductLaunchNotMarked ProductitemMarkedForProductLaunch = "not_marked"
)

func (value ProductitemMarkedForProductLaunch) String() string {
	return string(value)
}
