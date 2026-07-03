package objects

type AdMonetizationProperty struct {
	OwnerBusiness *Business `json:"owner_business,omitempty"`
}

var AdMonetizationPropertyFields = struct {
	OwnerBusiness string
}{
	OwnerBusiness: "owner_business",
}
