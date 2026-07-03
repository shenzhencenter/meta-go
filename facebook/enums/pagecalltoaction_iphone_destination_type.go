package enums

type PagecalltoactionIphoneDestinationType string

const (
	PagecalltoactionIphoneDestinationTypeAppDeeplink              PagecalltoactionIphoneDestinationType = "APP_DEEPLINK"
	PagecalltoactionIphoneDestinationTypeBecomeAVolunteer         PagecalltoactionIphoneDestinationType = "BECOME_A_VOLUNTEER"
	PagecalltoactionIphoneDestinationTypeEmail                    PagecalltoactionIphoneDestinationType = "EMAIL"
	PagecalltoactionIphoneDestinationTypeFacebookApp              PagecalltoactionIphoneDestinationType = "FACEBOOK_APP"
	PagecalltoactionIphoneDestinationTypeFollow                   PagecalltoactionIphoneDestinationType = "FOLLOW"
	PagecalltoactionIphoneDestinationTypeMarketplaceInventoryPage PagecalltoactionIphoneDestinationType = "MARKETPLACE_INVENTORY_PAGE"
	PagecalltoactionIphoneDestinationTypeMenuOnFacebook           PagecalltoactionIphoneDestinationType = "MENU_ON_FACEBOOK"
	PagecalltoactionIphoneDestinationTypeMessenger                PagecalltoactionIphoneDestinationType = "MESSENGER"
	PagecalltoactionIphoneDestinationTypeMiniShop                 PagecalltoactionIphoneDestinationType = "MINI_SHOP"
	PagecalltoactionIphoneDestinationTypeNone                     PagecalltoactionIphoneDestinationType = "NONE"
	PagecalltoactionIphoneDestinationTypePhoneCall                PagecalltoactionIphoneDestinationType = "PHONE_CALL"
	PagecalltoactionIphoneDestinationTypeShopOnFacebook           PagecalltoactionIphoneDestinationType = "SHOP_ON_FACEBOOK"
	PagecalltoactionIphoneDestinationTypeWebsite                  PagecalltoactionIphoneDestinationType = "WEBSITE"
)

func (value PagecalltoactionIphoneDestinationType) String() string {
	return string(value)
}
