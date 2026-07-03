package objects

type ProductItemImporterAddress struct {
	City       *string `json:"city,omitempty"`
	Country    *string `json:"country,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	Region     *string `json:"region,omitempty"`
	Street1    *string `json:"street1,omitempty"`
	Street2    *string `json:"street2,omitempty"`
}
