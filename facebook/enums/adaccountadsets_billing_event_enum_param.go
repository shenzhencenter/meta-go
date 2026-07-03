package enums

type AdaccountadsetsBillingEventEnumParam string

const (
	AdaccountadsetsBillingEventEnumParamAppInstalls        AdaccountadsetsBillingEventEnumParam = "APP_INSTALLS"
	AdaccountadsetsBillingEventEnumParamClicks             AdaccountadsetsBillingEventEnumParam = "CLICKS"
	AdaccountadsetsBillingEventEnumParamImpressions        AdaccountadsetsBillingEventEnumParam = "IMPRESSIONS"
	AdaccountadsetsBillingEventEnumParamLinkClicks         AdaccountadsetsBillingEventEnumParam = "LINK_CLICKS"
	AdaccountadsetsBillingEventEnumParamListingInteraction AdaccountadsetsBillingEventEnumParam = "LISTING_INTERACTION"
	AdaccountadsetsBillingEventEnumParamNone               AdaccountadsetsBillingEventEnumParam = "NONE"
	AdaccountadsetsBillingEventEnumParamOfferClaims        AdaccountadsetsBillingEventEnumParam = "OFFER_CLAIMS"
	AdaccountadsetsBillingEventEnumParamPageLikes          AdaccountadsetsBillingEventEnumParam = "PAGE_LIKES"
	AdaccountadsetsBillingEventEnumParamPostEngagement     AdaccountadsetsBillingEventEnumParam = "POST_ENGAGEMENT"
	AdaccountadsetsBillingEventEnumParamPurchase           AdaccountadsetsBillingEventEnumParam = "PURCHASE"
	AdaccountadsetsBillingEventEnumParamThruplay           AdaccountadsetsBillingEventEnumParam = "THRUPLAY"
)

func (value AdaccountadsetsBillingEventEnumParam) String() string {
	return string(value)
}
