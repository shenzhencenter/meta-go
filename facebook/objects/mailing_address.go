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
