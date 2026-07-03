package enums

type AdcampaignDestinationType string

const (
	AdcampaignDestinationTypeApp                                       AdcampaignDestinationType = "APP"
	AdcampaignDestinationTypeApplinksAutomatic                         AdcampaignDestinationType = "APPLINKS_AUTOMATIC"
	AdcampaignDestinationTypeFacebook                                  AdcampaignDestinationType = "FACEBOOK"
	AdcampaignDestinationTypeFacebookLive                              AdcampaignDestinationType = "FACEBOOK_LIVE"
	AdcampaignDestinationTypeFacebookPage                              AdcampaignDestinationType = "FACEBOOK_PAGE"
	AdcampaignDestinationTypeImagine                                   AdcampaignDestinationType = "IMAGINE"
	AdcampaignDestinationTypeInstagramDirect                           AdcampaignDestinationType = "INSTAGRAM_DIRECT"
	AdcampaignDestinationTypeInstagramLive                             AdcampaignDestinationType = "INSTAGRAM_LIVE"
	AdcampaignDestinationTypeInstagramProfile                          AdcampaignDestinationType = "INSTAGRAM_PROFILE"
	AdcampaignDestinationTypeInstagramProfileAndFacebookPage           AdcampaignDestinationType = "INSTAGRAM_PROFILE_AND_FACEBOOK_PAGE"
	AdcampaignDestinationTypeMessagingInstagramDirectMessenger         AdcampaignDestinationType = "MESSAGING_INSTAGRAM_DIRECT_MESSENGER"
	AdcampaignDestinationTypeMessagingInstagramDirectMessengerWhatsapp AdcampaignDestinationType = "MESSAGING_INSTAGRAM_DIRECT_MESSENGER_WHATSAPP"
	AdcampaignDestinationTypeMessagingInstagramDirectWhatsapp          AdcampaignDestinationType = "MESSAGING_INSTAGRAM_DIRECT_WHATSAPP"
	AdcampaignDestinationTypeMessagingMessengerWhatsapp                AdcampaignDestinationType = "MESSAGING_MESSENGER_WHATSAPP"
	AdcampaignDestinationTypeMessenger                                 AdcampaignDestinationType = "MESSENGER"
	AdcampaignDestinationTypeOnAd                                      AdcampaignDestinationType = "ON_AD"
	AdcampaignDestinationTypeOnEvent                                   AdcampaignDestinationType = "ON_EVENT"
	AdcampaignDestinationTypeOnPage                                    AdcampaignDestinationType = "ON_PAGE"
	AdcampaignDestinationTypeOnPost                                    AdcampaignDestinationType = "ON_POST"
	AdcampaignDestinationTypeOnVideo                                   AdcampaignDestinationType = "ON_VIDEO"
	AdcampaignDestinationTypeShopAutomatic                             AdcampaignDestinationType = "SHOP_AUTOMATIC"
	AdcampaignDestinationTypeWebsite                                   AdcampaignDestinationType = "WEBSITE"
	AdcampaignDestinationTypeWhatsapp                                  AdcampaignDestinationType = "WHATSAPP"
)

func (value AdcampaignDestinationType) String() string {
	return string(value)
}
