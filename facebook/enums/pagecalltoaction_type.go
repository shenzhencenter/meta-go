package enums

type PagecalltoactionType string

const (
	PagecalltoactionTypeBecomeAVolunteer   PagecalltoactionType = "BECOME_A_VOLUNTEER"
	PagecalltoactionTypeBookAppointment    PagecalltoactionType = "BOOK_APPOINTMENT"
	PagecalltoactionTypeBookNow            PagecalltoactionType = "BOOK_NOW"
	PagecalltoactionTypeBuyTickets         PagecalltoactionType = "BUY_TICKETS"
	PagecalltoactionTypeCallNow            PagecalltoactionType = "CALL_NOW"
	PagecalltoactionTypeCharityDonate      PagecalltoactionType = "CHARITY_DONATE"
	PagecalltoactionTypeCheckIn            PagecalltoactionType = "CHECK_IN"
	PagecalltoactionTypeContactUs          PagecalltoactionType = "CONTACT_US"
	PagecalltoactionTypeCreatorStorefront  PagecalltoactionType = "CREATOR_STOREFRONT"
	PagecalltoactionTypeDonateNow          PagecalltoactionType = "DONATE_NOW"
	PagecalltoactionTypeEmail              PagecalltoactionType = "EMAIL"
	PagecalltoactionTypeFollowPage         PagecalltoactionType = "FOLLOW_PAGE"
	PagecalltoactionTypeGetDirections      PagecalltoactionType = "GET_DIRECTIONS"
	PagecalltoactionTypeGetOffer           PagecalltoactionType = "GET_OFFER"
	PagecalltoactionTypeGetOfferView       PagecalltoactionType = "GET_OFFER_VIEW"
	PagecalltoactionTypeInterested         PagecalltoactionType = "INTERESTED"
	PagecalltoactionTypeLearnMore          PagecalltoactionType = "LEARN_MORE"
	PagecalltoactionTypeListen             PagecalltoactionType = "LISTEN"
	PagecalltoactionTypeLocalDevPlatform   PagecalltoactionType = "LOCAL_DEV_PLATFORM"
	PagecalltoactionTypeMessage            PagecalltoactionType = "MESSAGE"
	PagecalltoactionTypeMobileCenter       PagecalltoactionType = "MOBILE_CENTER"
	PagecalltoactionTypeOpenApp            PagecalltoactionType = "OPEN_APP"
	PagecalltoactionTypeOrderFood          PagecalltoactionType = "ORDER_FOOD"
	PagecalltoactionTypePlayMusic          PagecalltoactionType = "PLAY_MUSIC"
	PagecalltoactionTypePlayNow            PagecalltoactionType = "PLAY_NOW"
	PagecalltoactionTypePurchaseGiftCards  PagecalltoactionType = "PURCHASE_GIFT_CARDS"
	PagecalltoactionTypeRequestAppointment PagecalltoactionType = "REQUEST_APPOINTMENT"
	PagecalltoactionTypeRequestQuote       PagecalltoactionType = "REQUEST_QUOTE"
	PagecalltoactionTypeShopNow            PagecalltoactionType = "SHOP_NOW"
	PagecalltoactionTypeShopOnFacebook     PagecalltoactionType = "SHOP_ON_FACEBOOK"
	PagecalltoactionTypeSignUp             PagecalltoactionType = "SIGN_UP"
	PagecalltoactionTypeViewInventory      PagecalltoactionType = "VIEW_INVENTORY"
	PagecalltoactionTypeViewMenu           PagecalltoactionType = "VIEW_MENU"
	PagecalltoactionTypeViewShop           PagecalltoactionType = "VIEW_SHOP"
	PagecalltoactionTypeVisitGroup         PagecalltoactionType = "VISIT_GROUP"
	PagecalltoactionTypeWatchNow           PagecalltoactionType = "WATCH_NOW"
	PagecalltoactionTypeWoodhengeSupport   PagecalltoactionType = "WOODHENGE_SUPPORT"
)

func (value PagecalltoactionType) String() string {
	return string(value)
}
