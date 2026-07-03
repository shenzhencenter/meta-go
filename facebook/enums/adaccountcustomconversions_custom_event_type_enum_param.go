package enums

type AdaccountcustomconversionsCustomEventTypeEnumParam string

const (
	AdaccountcustomconversionsCustomEventTypeEnumParamAddPaymentInfo       AdaccountcustomconversionsCustomEventTypeEnumParam = "ADD_PAYMENT_INFO"
	AdaccountcustomconversionsCustomEventTypeEnumParamAddToCart            AdaccountcustomconversionsCustomEventTypeEnumParam = "ADD_TO_CART"
	AdaccountcustomconversionsCustomEventTypeEnumParamAddToWishlist        AdaccountcustomconversionsCustomEventTypeEnumParam = "ADD_TO_WISHLIST"
	AdaccountcustomconversionsCustomEventTypeEnumParamCompleteRegistration AdaccountcustomconversionsCustomEventTypeEnumParam = "COMPLETE_REGISTRATION"
	AdaccountcustomconversionsCustomEventTypeEnumParamContact              AdaccountcustomconversionsCustomEventTypeEnumParam = "CONTACT"
	AdaccountcustomconversionsCustomEventTypeEnumParamContentView          AdaccountcustomconversionsCustomEventTypeEnumParam = "CONTENT_VIEW"
	AdaccountcustomconversionsCustomEventTypeEnumParamCustomizeProduct     AdaccountcustomconversionsCustomEventTypeEnumParam = "CUSTOMIZE_PRODUCT"
	AdaccountcustomconversionsCustomEventTypeEnumParamDonate               AdaccountcustomconversionsCustomEventTypeEnumParam = "DONATE"
	AdaccountcustomconversionsCustomEventTypeEnumParamFacebookSelected     AdaccountcustomconversionsCustomEventTypeEnumParam = "FACEBOOK_SELECTED"
	AdaccountcustomconversionsCustomEventTypeEnumParamFindLocation         AdaccountcustomconversionsCustomEventTypeEnumParam = "FIND_LOCATION"
	AdaccountcustomconversionsCustomEventTypeEnumParamInitiatedCheckout    AdaccountcustomconversionsCustomEventTypeEnumParam = "INITIATED_CHECKOUT"
	AdaccountcustomconversionsCustomEventTypeEnumParamLead                 AdaccountcustomconversionsCustomEventTypeEnumParam = "LEAD"
	AdaccountcustomconversionsCustomEventTypeEnumParamListingInteraction   AdaccountcustomconversionsCustomEventTypeEnumParam = "LISTING_INTERACTION"
	AdaccountcustomconversionsCustomEventTypeEnumParamOther                AdaccountcustomconversionsCustomEventTypeEnumParam = "OTHER"
	AdaccountcustomconversionsCustomEventTypeEnumParamPurchase             AdaccountcustomconversionsCustomEventTypeEnumParam = "PURCHASE"
	AdaccountcustomconversionsCustomEventTypeEnumParamSchedule             AdaccountcustomconversionsCustomEventTypeEnumParam = "SCHEDULE"
	AdaccountcustomconversionsCustomEventTypeEnumParamSearch               AdaccountcustomconversionsCustomEventTypeEnumParam = "SEARCH"
	AdaccountcustomconversionsCustomEventTypeEnumParamStartTrial           AdaccountcustomconversionsCustomEventTypeEnumParam = "START_TRIAL"
	AdaccountcustomconversionsCustomEventTypeEnumParamSubmitApplication    AdaccountcustomconversionsCustomEventTypeEnumParam = "SUBMIT_APPLICATION"
	AdaccountcustomconversionsCustomEventTypeEnumParamSubscribe            AdaccountcustomconversionsCustomEventTypeEnumParam = "SUBSCRIBE"
)

func (value AdaccountcustomconversionsCustomEventTypeEnumParam) String() string {
	return string(value)
}
