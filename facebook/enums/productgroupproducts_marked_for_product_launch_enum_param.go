package enums

type ProductgroupproductsMarkedForProductLaunchEnumParam string

const (
	ProductgroupproductsMarkedForProductLaunchEnumParamDefault   ProductgroupproductsMarkedForProductLaunchEnumParam = "default"
	ProductgroupproductsMarkedForProductLaunchEnumParamMarked    ProductgroupproductsMarkedForProductLaunchEnumParam = "marked"
	ProductgroupproductsMarkedForProductLaunchEnumParamNotMarked ProductgroupproductsMarkedForProductLaunchEnumParam = "not_marked"
)

func (value ProductgroupproductsMarkedForProductLaunchEnumParam) String() string {
	return string(value)
}
