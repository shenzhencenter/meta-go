package enums

type AdaccountcustomaudiencesContentTypeEnumParam string

const (
	AdaccountcustomaudiencesContentTypeEnumParamAutomotiveModel      AdaccountcustomaudiencesContentTypeEnumParam = "AUTOMOTIVE_MODEL"
	AdaccountcustomaudiencesContentTypeEnumParamDestination          AdaccountcustomaudiencesContentTypeEnumParam = "DESTINATION"
	AdaccountcustomaudiencesContentTypeEnumParamFlight               AdaccountcustomaudiencesContentTypeEnumParam = "FLIGHT"
	AdaccountcustomaudiencesContentTypeEnumParamGeneric              AdaccountcustomaudiencesContentTypeEnumParam = "GENERIC"
	AdaccountcustomaudiencesContentTypeEnumParamHomeListing          AdaccountcustomaudiencesContentTypeEnumParam = "HOME_LISTING"
	AdaccountcustomaudiencesContentTypeEnumParamHotel                AdaccountcustomaudiencesContentTypeEnumParam = "HOTEL"
	AdaccountcustomaudiencesContentTypeEnumParamLocalServiceBusiness AdaccountcustomaudiencesContentTypeEnumParam = "LOCAL_SERVICE_BUSINESS"
	AdaccountcustomaudiencesContentTypeEnumParamMediaTitle           AdaccountcustomaudiencesContentTypeEnumParam = "MEDIA_TITLE"
	AdaccountcustomaudiencesContentTypeEnumParamOfflineProduct       AdaccountcustomaudiencesContentTypeEnumParam = "OFFLINE_PRODUCT"
	AdaccountcustomaudiencesContentTypeEnumParamProduct              AdaccountcustomaudiencesContentTypeEnumParam = "PRODUCT"
	AdaccountcustomaudiencesContentTypeEnumParamVehicle              AdaccountcustomaudiencesContentTypeEnumParam = "VEHICLE"
	AdaccountcustomaudiencesContentTypeEnumParamVehicleOffer         AdaccountcustomaudiencesContentTypeEnumParam = "VEHICLE_OFFER"
)

func (value AdaccountcustomaudiencesContentTypeEnumParam) String() string {
	return string(value)
}
