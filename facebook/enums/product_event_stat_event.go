package enums

type ProducteventstatEvent string

const (
	ProducteventstatEventAddtocart        ProducteventstatEvent = "AddToCart"
	ProducteventstatEventAddtowishlist    ProducteventstatEvent = "AddToWishlist"
	ProducteventstatEventInitiatecheckout ProducteventstatEvent = "InitiateCheckout"
	ProducteventstatEventLead             ProducteventstatEvent = "Lead"
	ProducteventstatEventPurchase         ProducteventstatEvent = "Purchase"
	ProducteventstatEventSearch           ProducteventstatEvent = "Search"
	ProducteventstatEventSubscribe        ProducteventstatEvent = "Subscribe"
	ProducteventstatEventViewcontent      ProducteventstatEvent = "ViewContent"
)

func (value ProducteventstatEvent) String() string {
	return string(value)
}
