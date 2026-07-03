package enums

type PagecalltoactionWebDestinationType string

const (
	PagecalltoactionWebDestinationTypeBecomeAVolunteer PagecalltoactionWebDestinationType = "BECOME_A_VOLUNTEER"
	PagecalltoactionWebDestinationTypeBecomeSupporter  PagecalltoactionWebDestinationType = "BECOME_SUPPORTER"
	PagecalltoactionWebDestinationTypeEmail            PagecalltoactionWebDestinationType = "EMAIL"
	PagecalltoactionWebDestinationTypeFollow           PagecalltoactionWebDestinationType = "FOLLOW"
	PagecalltoactionWebDestinationTypeMessenger        PagecalltoactionWebDestinationType = "MESSENGER"
	PagecalltoactionWebDestinationTypeMobileCenter     PagecalltoactionWebDestinationType = "MOBILE_CENTER"
	PagecalltoactionWebDestinationTypeNone             PagecalltoactionWebDestinationType = "NONE"
	PagecalltoactionWebDestinationTypeShopOnFacebook   PagecalltoactionWebDestinationType = "SHOP_ON_FACEBOOK"
	PagecalltoactionWebDestinationTypeWebsite          PagecalltoactionWebDestinationType = "WEBSITE"
)

func (value PagecalltoactionWebDestinationType) String() string {
	return string(value)
}
