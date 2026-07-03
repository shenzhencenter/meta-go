package objects

type ProductItemImporterAddress struct {
	City       *string `json:"city,omitempty"`
	Country    *string `json:"country,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	Region     *string `json:"region,omitempty"`
	Street1    *string `json:"street1,omitempty"`
	Street2    *string `json:"street2,omitempty"`
}

var ProductItemImporterAddressFields = struct {
	City       string
	Country    string
	PostalCode string
	Region     string
	Street1    string
	Street2    string
}{
	City:       "city",
	Country:    "country",
	PostalCode: "postal_code",
	Region:     "region",
	Street1:    "street1",
	Street2:    "street2",
}
