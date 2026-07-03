package enums

type AdcreativevideodatacustomoverlayspecOption string

const (
	AdcreativevideodatacustomoverlayspecOptionBankTransfer       AdcreativevideodatacustomoverlayspecOption = "bank_transfer"
	AdcreativevideodatacustomoverlayspecOptionBoleto             AdcreativevideodatacustomoverlayspecOption = "boleto"
	AdcreativevideodatacustomoverlayspecOptionCashOnDelivery     AdcreativevideodatacustomoverlayspecOption = "cash_on_delivery"
	AdcreativevideodatacustomoverlayspecOptionDiscountWithBoleto AdcreativevideodatacustomoverlayspecOption = "discount_with_boleto"
	AdcreativevideodatacustomoverlayspecOptionFastDelivery       AdcreativevideodatacustomoverlayspecOption = "fast_delivery"
	AdcreativevideodatacustomoverlayspecOptionFreeShipping       AdcreativevideodatacustomoverlayspecOption = "free_shipping"
	AdcreativevideodatacustomoverlayspecOptionHomeDelivery       AdcreativevideodatacustomoverlayspecOption = "home_delivery"
	AdcreativevideodatacustomoverlayspecOptionInventory          AdcreativevideodatacustomoverlayspecOption = "inventory"
	AdcreativevideodatacustomoverlayspecOptionPayAtHotel         AdcreativevideodatacustomoverlayspecOption = "pay_at_hotel"
	AdcreativevideodatacustomoverlayspecOptionPayOnArrival       AdcreativevideodatacustomoverlayspecOption = "pay_on_arrival"
)

func (value AdcreativevideodatacustomoverlayspecOption) String() string {
	return string(value)
}
