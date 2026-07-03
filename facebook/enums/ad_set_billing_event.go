package enums

type AdsetBillingEvent string

const (
	AdsetBillingEventAppInstalls        AdsetBillingEvent = "APP_INSTALLS"
	AdsetBillingEventClicks             AdsetBillingEvent = "CLICKS"
	AdsetBillingEventImpressions        AdsetBillingEvent = "IMPRESSIONS"
	AdsetBillingEventLinkClicks         AdsetBillingEvent = "LINK_CLICKS"
	AdsetBillingEventListingInteraction AdsetBillingEvent = "LISTING_INTERACTION"
	AdsetBillingEventNone               AdsetBillingEvent = "NONE"
	AdsetBillingEventOfferClaims        AdsetBillingEvent = "OFFER_CLAIMS"
	AdsetBillingEventPageLikes          AdsetBillingEvent = "PAGE_LIKES"
	AdsetBillingEventPostEngagement     AdsetBillingEvent = "POST_ENGAGEMENT"
	AdsetBillingEventPurchase           AdsetBillingEvent = "PURCHASE"
	AdsetBillingEventThruplay           AdsetBillingEvent = "THRUPLAY"
)

func (value AdsetBillingEvent) String() string {
	return string(value)
}
