package enums

type PagelocationsPickupOptionsEnumParam string

const (
	PagelocationsPickupOptionsEnumParamCurbside PagelocationsPickupOptionsEnumParam = "CURBSIDE"
	PagelocationsPickupOptionsEnumParamInStore  PagelocationsPickupOptionsEnumParam = "IN_STORE"
	PagelocationsPickupOptionsEnumParamOther    PagelocationsPickupOptionsEnumParam = "OTHER"
)

func (value PagelocationsPickupOptionsEnumParam) String() string {
	return string(value)
}
