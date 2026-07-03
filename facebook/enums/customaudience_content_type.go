package enums

type CustomaudienceContentType string

const (
	CustomaudienceContentTypeAutomotiveModel      CustomaudienceContentType = "AUTOMOTIVE_MODEL"
	CustomaudienceContentTypeDestination          CustomaudienceContentType = "DESTINATION"
	CustomaudienceContentTypeFlight               CustomaudienceContentType = "FLIGHT"
	CustomaudienceContentTypeGeneric              CustomaudienceContentType = "GENERIC"
	CustomaudienceContentTypeHomeListing          CustomaudienceContentType = "HOME_LISTING"
	CustomaudienceContentTypeHotel                CustomaudienceContentType = "HOTEL"
	CustomaudienceContentTypeLocalServiceBusiness CustomaudienceContentType = "LOCAL_SERVICE_BUSINESS"
	CustomaudienceContentTypeMediaTitle           CustomaudienceContentType = "MEDIA_TITLE"
	CustomaudienceContentTypeOfflineProduct       CustomaudienceContentType = "OFFLINE_PRODUCT"
	CustomaudienceContentTypeProduct              CustomaudienceContentType = "PRODUCT"
	CustomaudienceContentTypeVehicle              CustomaudienceContentType = "VEHICLE"
	CustomaudienceContentTypeVehicleOffer         CustomaudienceContentType = "VEHICLE_OFFER"
)

func (value CustomaudienceContentType) String() string {
	return string(value)
}
