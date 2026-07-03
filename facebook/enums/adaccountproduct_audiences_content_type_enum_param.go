package enums

type AdaccountproductAudiencesContentTypeEnumParam string

const (
	AdaccountproductAudiencesContentTypeEnumParamAutomotiveModel      AdaccountproductAudiencesContentTypeEnumParam = "AUTOMOTIVE_MODEL"
	AdaccountproductAudiencesContentTypeEnumParamDestination          AdaccountproductAudiencesContentTypeEnumParam = "DESTINATION"
	AdaccountproductAudiencesContentTypeEnumParamFlight               AdaccountproductAudiencesContentTypeEnumParam = "FLIGHT"
	AdaccountproductAudiencesContentTypeEnumParamGeneric              AdaccountproductAudiencesContentTypeEnumParam = "GENERIC"
	AdaccountproductAudiencesContentTypeEnumParamHomeListing          AdaccountproductAudiencesContentTypeEnumParam = "HOME_LISTING"
	AdaccountproductAudiencesContentTypeEnumParamHotel                AdaccountproductAudiencesContentTypeEnumParam = "HOTEL"
	AdaccountproductAudiencesContentTypeEnumParamLocalServiceBusiness AdaccountproductAudiencesContentTypeEnumParam = "LOCAL_SERVICE_BUSINESS"
	AdaccountproductAudiencesContentTypeEnumParamMediaTitle           AdaccountproductAudiencesContentTypeEnumParam = "MEDIA_TITLE"
	AdaccountproductAudiencesContentTypeEnumParamOfflineProduct       AdaccountproductAudiencesContentTypeEnumParam = "OFFLINE_PRODUCT"
	AdaccountproductAudiencesContentTypeEnumParamProduct              AdaccountproductAudiencesContentTypeEnumParam = "PRODUCT"
	AdaccountproductAudiencesContentTypeEnumParamVehicle              AdaccountproductAudiencesContentTypeEnumParam = "VEHICLE"
	AdaccountproductAudiencesContentTypeEnumParamVehicleOffer         AdaccountproductAudiencesContentTypeEnumParam = "VEHICLE_OFFER"
)

func (value AdaccountproductAudiencesContentTypeEnumParam) String() string {
	return string(value)
}
