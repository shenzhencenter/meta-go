package enums

type AdgrouppreviewsRenderTypeEnumParam string

const (
	AdgrouppreviewsRenderTypeEnumParamFallback AdgrouppreviewsRenderTypeEnumParam = "FALLBACK"
)

func (value AdgrouppreviewsRenderTypeEnumParam) String() string {
	return string(value)
}
