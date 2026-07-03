package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type MailingAddress struct {
	City       *string  `json:"city,omitempty"`
	CityPage   *Page    `json:"city_page,omitempty"`
	Country    *string  `json:"country,omitempty"`
	ID         *core.ID `json:"id,omitempty"`
	PostalCode *string  `json:"postal_code,omitempty"`
	Region     *string  `json:"region,omitempty"`
	Street1    *string  `json:"street1,omitempty"`
	Street2    *string  `json:"street2,omitempty"`
}

var MailingAddressFields = struct {
	City       string
	CityPage   string
	Country    string
	ID         string
	PostalCode string
	Region     string
	Street1    string
	Street2    string
}{
	City:       "city",
	CityPage:   "city_page",
	Country:    "country",
	ID:         "id",
	PostalCode: "postal_code",
	Region:     "region",
	Street1:    "street1",
	Street2:    "street2",
}
