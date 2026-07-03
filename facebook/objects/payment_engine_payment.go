package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type PaymentEnginePayment struct {
	Actions                   *[]map[string]interface{} `json:"actions,omitempty"`
	Application               *Application              `json:"application,omitempty"`
	Country                   *string                   `json:"country,omitempty"`
	CreatedTime               *time.Time                `json:"created_time,omitempty"`
	Disputes                  *[]map[string]interface{} `json:"disputes,omitempty"`
	FraudStatus               *string                   `json:"fraud_status,omitempty"`
	FulfillmentStatus         *string                   `json:"fulfillment_status,omitempty"`
	ID                        *core.ID                  `json:"id,omitempty"`
	IsFromAd                  *bool                     `json:"is_from_ad,omitempty"`
	IsFromPagePost            *bool                     `json:"is_from_page_post,omitempty"`
	Items                     *[]map[string]interface{} `json:"items,omitempty"`
	PayoutForeignExchangeRate *float64                  `json:"payout_foreign_exchange_rate,omitempty"`
	PhoneSupportEligible      *bool                     `json:"phone_support_eligible,omitempty"`
	Platform                  *string                   `json:"platform,omitempty"`
	RefundableAmount          *CurrencyAmount           `json:"refundable_amount,omitempty"`
	RequestID                 *core.ID                  `json:"request_id,omitempty"`
	Tax                       *string                   `json:"tax,omitempty"`
	TaxCountry                *string                   `json:"tax_country,omitempty"`
	Test                      *uint64                   `json:"test,omitempty"`
	User                      *User                     `json:"user,omitempty"`
}
