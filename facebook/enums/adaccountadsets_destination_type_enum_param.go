package enums

type AdaccountadsetsDestinationTypeEnumParam string

const (
	AdaccountadsetsDestinationTypeEnumParamApp                                       AdaccountadsetsDestinationTypeEnumParam = "APP"
	AdaccountadsetsDestinationTypeEnumParamApplinksAutomatic                         AdaccountadsetsDestinationTypeEnumParam = "APPLINKS_AUTOMATIC"
	AdaccountadsetsDestinationTypeEnumParamFacebook                                  AdaccountadsetsDestinationTypeEnumParam = "FACEBOOK"
	AdaccountadsetsDestinationTypeEnumParamFacebookLive                              AdaccountadsetsDestinationTypeEnumParam = "FACEBOOK_LIVE"
	AdaccountadsetsDestinationTypeEnumParamFacebookPage                              AdaccountadsetsDestinationTypeEnumParam = "FACEBOOK_PAGE"
	AdaccountadsetsDestinationTypeEnumParamImagine                                   AdaccountadsetsDestinationTypeEnumParam = "IMAGINE"
	AdaccountadsetsDestinationTypeEnumParamInstagramDirect                           AdaccountadsetsDestinationTypeEnumParam = "INSTAGRAM_DIRECT"
	AdaccountadsetsDestinationTypeEnumParamInstagramLive                             AdaccountadsetsDestinationTypeEnumParam = "INSTAGRAM_LIVE"
	AdaccountadsetsDestinationTypeEnumParamInstagramProfile                          AdaccountadsetsDestinationTypeEnumParam = "INSTAGRAM_PROFILE"
	AdaccountadsetsDestinationTypeEnumParamInstagramProfileAndFacebookPage           AdaccountadsetsDestinationTypeEnumParam = "INSTAGRAM_PROFILE_AND_FACEBOOK_PAGE"
	AdaccountadsetsDestinationTypeEnumParamMessagingInstagramDirectMessenger         AdaccountadsetsDestinationTypeEnumParam = "MESSAGING_INSTAGRAM_DIRECT_MESSENGER"
	AdaccountadsetsDestinationTypeEnumParamMessagingInstagramDirectMessengerWhatsapp AdaccountadsetsDestinationTypeEnumParam = "MESSAGING_INSTAGRAM_DIRECT_MESSENGER_WHATSAPP"
	AdaccountadsetsDestinationTypeEnumParamMessagingInstagramDirectWhatsapp          AdaccountadsetsDestinationTypeEnumParam = "MESSAGING_INSTAGRAM_DIRECT_WHATSAPP"
	AdaccountadsetsDestinationTypeEnumParamMessagingMessengerWhatsapp                AdaccountadsetsDestinationTypeEnumParam = "MESSAGING_MESSENGER_WHATSAPP"
	AdaccountadsetsDestinationTypeEnumParamMessenger                                 AdaccountadsetsDestinationTypeEnumParam = "MESSENGER"
	AdaccountadsetsDestinationTypeEnumParamOnAd                                      AdaccountadsetsDestinationTypeEnumParam = "ON_AD"
	AdaccountadsetsDestinationTypeEnumParamOnEvent                                   AdaccountadsetsDestinationTypeEnumParam = "ON_EVENT"
	AdaccountadsetsDestinationTypeEnumParamOnPage                                    AdaccountadsetsDestinationTypeEnumParam = "ON_PAGE"
	AdaccountadsetsDestinationTypeEnumParamOnPost                                    AdaccountadsetsDestinationTypeEnumParam = "ON_POST"
	AdaccountadsetsDestinationTypeEnumParamOnVideo                                   AdaccountadsetsDestinationTypeEnumParam = "ON_VIDEO"
	AdaccountadsetsDestinationTypeEnumParamShopAutomatic                             AdaccountadsetsDestinationTypeEnumParam = "SHOP_AUTOMATIC"
	AdaccountadsetsDestinationTypeEnumParamWebsite                                   AdaccountadsetsDestinationTypeEnumParam = "WEBSITE"
	AdaccountadsetsDestinationTypeEnumParamWhatsapp                                  AdaccountadsetsDestinationTypeEnumParam = "WHATSAPP"
)

func (value AdaccountadsetsDestinationTypeEnumParam) String() string {
	return string(value)
}
