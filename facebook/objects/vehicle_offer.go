package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type VehicleOffer struct {
	AmountCurrency       *string                             `json:"amount_currency,omitempty"`
	AmountPercentage     *float64                            `json:"amount_percentage,omitempty"`
	AmountPrice          *string                             `json:"amount_price,omitempty"`
	AmountQualifier      *string                             `json:"amount_qualifier,omitempty"`
	Applinks             *CatalogItemAppLinks                `json:"applinks,omitempty"`
	Availability         *string                             `json:"availability,omitempty"`
	BodyStyle            *string                             `json:"body_style,omitempty"`
	CashbackCurrency     *string                             `json:"cashback_currency,omitempty"`
	CashbackPrice        *string                             `json:"cashback_price,omitempty"`
	ComscoreMarketCodes  *[]string                           `json:"comscore_market_codes,omitempty"`
	Currency             *string                             `json:"currency,omitempty"`
	CustomLabelX0        *string                             `json:"custom_label_0,omitempty"`
	CustomLabelX1        *string                             `json:"custom_label_1,omitempty"`
	CustomLabelX2        *string                             `json:"custom_label_2,omitempty"`
	CustomLabelX3        *string                             `json:"custom_label_3,omitempty"`
	CustomLabelX4        *string                             `json:"custom_label_4,omitempty"`
	CustomNumberX0       *uint64                             `json:"custom_number_0,omitempty"`
	CustomNumberX1       *uint64                             `json:"custom_number_1,omitempty"`
	CustomNumberX2       *uint64                             `json:"custom_number_2,omitempty"`
	CustomNumberX3       *uint64                             `json:"custom_number_3,omitempty"`
	CustomNumberX4       *uint64                             `json:"custom_number_4,omitempty"`
	DmaCodes             *[]string                           `json:"dma_codes,omitempty"`
	DownpaymentCurrency  *string                             `json:"downpayment_currency,omitempty"`
	DownpaymentPrice     *string                             `json:"downpayment_price,omitempty"`
	DownpaymentQualifier *string                             `json:"downpayment_qualifier,omitempty"`
	Drivetrain           *string                             `json:"drivetrain,omitempty"`
	EndDate              *string                             `json:"end_date,omitempty"`
	EndTime              *int                                `json:"end_time,omitempty"`
	ExteriorColor        *string                             `json:"exterior_color,omitempty"`
	FuelType             *string                             `json:"fuel_type,omitempty"`
	Generation           *string                             `json:"generation,omitempty"`
	ID                   *core.ID                            `json:"id,omitempty"`
	ImageFetchStatus     *enums.VehicleofferImageFetchStatus `json:"image_fetch_status,omitempty"`
	Images               *[]string                           `json:"images,omitempty"`
	InteriorColor        *string                             `json:"interior_color,omitempty"`
	InteriorUpholstery   *string                             `json:"interior_upholstery,omitempty"`
	Make                 *string                             `json:"make,omitempty"`
	Model                *string                             `json:"model,omitempty"`
	OfferDescription     *string                             `json:"offer_description,omitempty"`
	OfferDisclaimer      *string                             `json:"offer_disclaimer,omitempty"`
	OfferType            *string                             `json:"offer_type,omitempty"`
	Price                *string                             `json:"price,omitempty"`
	ProductPriorityX0    *float64                            `json:"product_priority_0,omitempty"`
	ProductPriorityX1    *float64                            `json:"product_priority_1,omitempty"`
	ProductPriorityX2    *float64                            `json:"product_priority_2,omitempty"`
	ProductPriorityX3    *float64                            `json:"product_priority_3,omitempty"`
	ProductPriorityX4    *float64                            `json:"product_priority_4,omitempty"`
	SanitizedImages      *[]string                           `json:"sanitized_images,omitempty"`
	StartDate            *string                             `json:"start_date,omitempty"`
	StartTime            *int                                `json:"start_time,omitempty"`
	Tags                 *[]string                           `json:"tags,omitempty"`
	TermLength           *uint64                             `json:"term_length,omitempty"`
	TermQualifier        *string                             `json:"term_qualifier,omitempty"`
	Title                *string                             `json:"title,omitempty"`
	Transmission         *string                             `json:"transmission,omitempty"`
	Trim                 *string                             `json:"trim,omitempty"`
	UnitPrice            *map[string]interface{}             `json:"unit_price,omitempty"`
	URL                  *string                             `json:"url,omitempty"`
	VehicleOfferID       *core.ID                            `json:"vehicle_offer_id,omitempty"`
	Visibility           *enums.VehicleofferVisibility       `json:"visibility,omitempty"`
	Year                 *int                                `json:"year,omitempty"`
}
