package enums

type PageAttire string

const (
	PageAttireCasual      PageAttire = "Casual"
	PageAttireDressy      PageAttire = "Dressy"
	PageAttireUnspecified PageAttire = "Unspecified"
)

func (value PageAttire) String() string {
	return string(value)
}
