package enums

type ApplicationactivitiesEventEnumParam string

const (
	ApplicationactivitiesEventEnumParamCustomAppEvents  ApplicationactivitiesEventEnumParam = "CUSTOM_APP_EVENTS"
	ApplicationactivitiesEventEnumParamDeferredAppLink  ApplicationactivitiesEventEnumParam = "DEFERRED_APP_LINK"
	ApplicationactivitiesEventEnumParamMobileAppInstall ApplicationactivitiesEventEnumParam = "MOBILE_APP_INSTALL"
)

func (value ApplicationactivitiesEventEnumParam) String() string {
	return string(value)
}
