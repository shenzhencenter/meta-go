package enums

type AdcreativelinkdatacustomoverlayspecOption string

const (
	AdcreativelinkdatacustomoverlayspecOptionBankTransfer       AdcreativelinkdatacustomoverlayspecOption = "bank_transfer"
	AdcreativelinkdatacustomoverlayspecOptionBoleto             AdcreativelinkdatacustomoverlayspecOption = "boleto"
	AdcreativelinkdatacustomoverlayspecOptionCashOnDelivery     AdcreativelinkdatacustomoverlayspecOption = "cash_on_delivery"
	AdcreativelinkdatacustomoverlayspecOptionDiscountWithBoleto AdcreativelinkdatacustomoverlayspecOption = "discount_with_boleto"
	AdcreativelinkdatacustomoverlayspecOptionFastDelivery       AdcreativelinkdatacustomoverlayspecOption = "fast_delivery"
	AdcreativelinkdatacustomoverlayspecOptionFreeShipping       AdcreativelinkdatacustomoverlayspecOption = "free_shipping"
	AdcreativelinkdatacustomoverlayspecOptionHomeDelivery       AdcreativelinkdatacustomoverlayspecOption = "home_delivery"
	AdcreativelinkdatacustomoverlayspecOptionInventory          AdcreativelinkdatacustomoverlayspecOption = "inventory"
	AdcreativelinkdatacustomoverlayspecOptionPayAtHotel         AdcreativelinkdatacustomoverlayspecOption = "pay_at_hotel"
	AdcreativelinkdatacustomoverlayspecOptionPayOnArrival       AdcreativelinkdatacustomoverlayspecOption = "pay_on_arrival"
)

func (value AdcreativelinkdatacustomoverlayspecOption) String() string {
	return string(value)
}
