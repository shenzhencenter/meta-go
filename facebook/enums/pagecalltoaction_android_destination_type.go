package enums

type PagecalltoactionAndroidDestinationType string

const (
	PagecalltoactionAndroidDestinationTypeAppDeeplink              PagecalltoactionAndroidDestinationType = "APP_DEEPLINK"
	PagecalltoactionAndroidDestinationTypeBecomeAVolunteer         PagecalltoactionAndroidDestinationType = "BECOME_A_VOLUNTEER"
	PagecalltoactionAndroidDestinationTypeEmail                    PagecalltoactionAndroidDestinationType = "EMAIL"
	PagecalltoactionAndroidDestinationTypeFacebookApp              PagecalltoactionAndroidDestinationType = "FACEBOOK_APP"
	PagecalltoactionAndroidDestinationTypeFollow                   PagecalltoactionAndroidDestinationType = "FOLLOW"
	PagecalltoactionAndroidDestinationTypeMarketplaceInventoryPage PagecalltoactionAndroidDestinationType = "MARKETPLACE_INVENTORY_PAGE"
	PagecalltoactionAndroidDestinationTypeMenuOnFacebook           PagecalltoactionAndroidDestinationType = "MENU_ON_FACEBOOK"
	PagecalltoactionAndroidDestinationTypeMessenger                PagecalltoactionAndroidDestinationType = "MESSENGER"
	PagecalltoactionAndroidDestinationTypeMiniShop                 PagecalltoactionAndroidDestinationType = "MINI_SHOP"
	PagecalltoactionAndroidDestinationTypeMobileCenter             PagecalltoactionAndroidDestinationType = "MOBILE_CENTER"
	PagecalltoactionAndroidDestinationTypeNone                     PagecalltoactionAndroidDestinationType = "NONE"
	PagecalltoactionAndroidDestinationTypePhoneCall                PagecalltoactionAndroidDestinationType = "PHONE_CALL"
	PagecalltoactionAndroidDestinationTypeShopOnFacebook           PagecalltoactionAndroidDestinationType = "SHOP_ON_FACEBOOK"
	PagecalltoactionAndroidDestinationTypeWebsite                  PagecalltoactionAndroidDestinationType = "WEBSITE"
)

func (value PagecalltoactionAndroidDestinationType) String() string {
	return string(value)
}
