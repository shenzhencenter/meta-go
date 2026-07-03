package enums

type ProductgroupproductsVisibilityEnumParam string

const (
	ProductgroupproductsVisibilityEnumParamPublished ProductgroupproductsVisibilityEnumParam = "published"
	ProductgroupproductsVisibilityEnumParamStaging   ProductgroupproductsVisibilityEnumParam = "staging"
)

func (value ProductgroupproductsVisibilityEnumParam) String() string {
	return string(value)
}
