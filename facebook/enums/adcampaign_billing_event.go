package enums

type AdcampaignBillingEvent string

const (
	AdcampaignBillingEventAppInstalls        AdcampaignBillingEvent = "APP_INSTALLS"
	AdcampaignBillingEventClicks             AdcampaignBillingEvent = "CLICKS"
	AdcampaignBillingEventImpressions        AdcampaignBillingEvent = "IMPRESSIONS"
	AdcampaignBillingEventLinkClicks         AdcampaignBillingEvent = "LINK_CLICKS"
	AdcampaignBillingEventListingInteraction AdcampaignBillingEvent = "LISTING_INTERACTION"
	AdcampaignBillingEventNone               AdcampaignBillingEvent = "NONE"
	AdcampaignBillingEventOfferClaims        AdcampaignBillingEvent = "OFFER_CLAIMS"
	AdcampaignBillingEventPageLikes          AdcampaignBillingEvent = "PAGE_LIKES"
	AdcampaignBillingEventPostEngagement     AdcampaignBillingEvent = "POST_ENGAGEMENT"
	AdcampaignBillingEventPurchase           AdcampaignBillingEvent = "PURCHASE"
	AdcampaignBillingEventThruplay           AdcampaignBillingEvent = "THRUPLAY"
)

func (value AdcampaignBillingEvent) String() string {
	return string(value)
}
