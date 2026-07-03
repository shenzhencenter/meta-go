package enums

type BusinesscustomconversionsCustomEventTypeEnumParam string

const (
	BusinesscustomconversionsCustomEventTypeEnumParamAddPaymentInfo       BusinesscustomconversionsCustomEventTypeEnumParam = "ADD_PAYMENT_INFO"
	BusinesscustomconversionsCustomEventTypeEnumParamAddToCart            BusinesscustomconversionsCustomEventTypeEnumParam = "ADD_TO_CART"
	BusinesscustomconversionsCustomEventTypeEnumParamAddToWishlist        BusinesscustomconversionsCustomEventTypeEnumParam = "ADD_TO_WISHLIST"
	BusinesscustomconversionsCustomEventTypeEnumParamCompleteRegistration BusinesscustomconversionsCustomEventTypeEnumParam = "COMPLETE_REGISTRATION"
	BusinesscustomconversionsCustomEventTypeEnumParamContact              BusinesscustomconversionsCustomEventTypeEnumParam = "CONTACT"
	BusinesscustomconversionsCustomEventTypeEnumParamContentView          BusinesscustomconversionsCustomEventTypeEnumParam = "CONTENT_VIEW"
	BusinesscustomconversionsCustomEventTypeEnumParamCustomizeProduct     BusinesscustomconversionsCustomEventTypeEnumParam = "CUSTOMIZE_PRODUCT"
	BusinesscustomconversionsCustomEventTypeEnumParamDonate               BusinesscustomconversionsCustomEventTypeEnumParam = "DONATE"
	BusinesscustomconversionsCustomEventTypeEnumParamFacebookSelected     BusinesscustomconversionsCustomEventTypeEnumParam = "FACEBOOK_SELECTED"
	BusinesscustomconversionsCustomEventTypeEnumParamFindLocation         BusinesscustomconversionsCustomEventTypeEnumParam = "FIND_LOCATION"
	BusinesscustomconversionsCustomEventTypeEnumParamInitiatedCheckout    BusinesscustomconversionsCustomEventTypeEnumParam = "INITIATED_CHECKOUT"
	BusinesscustomconversionsCustomEventTypeEnumParamLead                 BusinesscustomconversionsCustomEventTypeEnumParam = "LEAD"
	BusinesscustomconversionsCustomEventTypeEnumParamListingInteraction   BusinesscustomconversionsCustomEventTypeEnumParam = "LISTING_INTERACTION"
	BusinesscustomconversionsCustomEventTypeEnumParamOther                BusinesscustomconversionsCustomEventTypeEnumParam = "OTHER"
	BusinesscustomconversionsCustomEventTypeEnumParamPurchase             BusinesscustomconversionsCustomEventTypeEnumParam = "PURCHASE"
	BusinesscustomconversionsCustomEventTypeEnumParamSchedule             BusinesscustomconversionsCustomEventTypeEnumParam = "SCHEDULE"
	BusinesscustomconversionsCustomEventTypeEnumParamSearch               BusinesscustomconversionsCustomEventTypeEnumParam = "SEARCH"
	BusinesscustomconversionsCustomEventTypeEnumParamStartTrial           BusinesscustomconversionsCustomEventTypeEnumParam = "START_TRIAL"
	BusinesscustomconversionsCustomEventTypeEnumParamSubmitApplication    BusinesscustomconversionsCustomEventTypeEnumParam = "SUBMIT_APPLICATION"
	BusinesscustomconversionsCustomEventTypeEnumParamSubscribe            BusinesscustomconversionsCustomEventTypeEnumParam = "SUBSCRIBE"
)

func (value BusinesscustomconversionsCustomEventTypeEnumParam) String() string {
	return string(value)
}
