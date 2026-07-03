package enums

type CustomconversionCustomEventType string

const (
	CustomconversionCustomEventTypeAddPaymentInfo       CustomconversionCustomEventType = "ADD_PAYMENT_INFO"
	CustomconversionCustomEventTypeAddToCart            CustomconversionCustomEventType = "ADD_TO_CART"
	CustomconversionCustomEventTypeAddToWishlist        CustomconversionCustomEventType = "ADD_TO_WISHLIST"
	CustomconversionCustomEventTypeCompleteRegistration CustomconversionCustomEventType = "COMPLETE_REGISTRATION"
	CustomconversionCustomEventTypeContact              CustomconversionCustomEventType = "CONTACT"
	CustomconversionCustomEventTypeContentView          CustomconversionCustomEventType = "CONTENT_VIEW"
	CustomconversionCustomEventTypeCustomizeProduct     CustomconversionCustomEventType = "CUSTOMIZE_PRODUCT"
	CustomconversionCustomEventTypeDonate               CustomconversionCustomEventType = "DONATE"
	CustomconversionCustomEventTypeFacebookSelected     CustomconversionCustomEventType = "FACEBOOK_SELECTED"
	CustomconversionCustomEventTypeFindLocation         CustomconversionCustomEventType = "FIND_LOCATION"
	CustomconversionCustomEventTypeInitiatedCheckout    CustomconversionCustomEventType = "INITIATED_CHECKOUT"
	CustomconversionCustomEventTypeLead                 CustomconversionCustomEventType = "LEAD"
	CustomconversionCustomEventTypeListingInteraction   CustomconversionCustomEventType = "LISTING_INTERACTION"
	CustomconversionCustomEventTypeOther                CustomconversionCustomEventType = "OTHER"
	CustomconversionCustomEventTypePurchase             CustomconversionCustomEventType = "PURCHASE"
	CustomconversionCustomEventTypeSchedule             CustomconversionCustomEventType = "SCHEDULE"
	CustomconversionCustomEventTypeSearch               CustomconversionCustomEventType = "SEARCH"
	CustomconversionCustomEventTypeStartTrial           CustomconversionCustomEventType = "START_TRIAL"
	CustomconversionCustomEventTypeSubmitApplication    CustomconversionCustomEventType = "SUBMIT_APPLICATION"
	CustomconversionCustomEventTypeSubscribe            CustomconversionCustomEventType = "SUBSCRIBE"
)

func (value CustomconversionCustomEventType) String() string {
	return string(value)
}
