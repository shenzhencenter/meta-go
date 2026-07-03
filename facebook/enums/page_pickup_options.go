package enums

type PagePickupOptions string

const (
	PagePickupOptionsCurbside PagePickupOptions = "CURBSIDE"
	PagePickupOptionsInStore  PagePickupOptions = "IN_STORE"
	PagePickupOptionsOther    PagePickupOptions = "OTHER"
)

func (value PagePickupOptions) String() string {
	return string(value)
}
