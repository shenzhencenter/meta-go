package enums

type AdaccountcustomaudiencesClaimObjectiveEnumParam string

const (
	AdaccountcustomaudiencesClaimObjectiveEnumParamAutomotiveModel  AdaccountcustomaudiencesClaimObjectiveEnumParam = "AUTOMOTIVE_MODEL"
	AdaccountcustomaudiencesClaimObjectiveEnumParamCollaborativeAds AdaccountcustomaudiencesClaimObjectiveEnumParam = "COLLABORATIVE_ADS"
	AdaccountcustomaudiencesClaimObjectiveEnumParamHomeListing      AdaccountcustomaudiencesClaimObjectiveEnumParam = "HOME_LISTING"
	AdaccountcustomaudiencesClaimObjectiveEnumParamMediaTitle       AdaccountcustomaudiencesClaimObjectiveEnumParam = "MEDIA_TITLE"
	AdaccountcustomaudiencesClaimObjectiveEnumParamProduct          AdaccountcustomaudiencesClaimObjectiveEnumParam = "PRODUCT"
	AdaccountcustomaudiencesClaimObjectiveEnumParamTravel           AdaccountcustomaudiencesClaimObjectiveEnumParam = "TRAVEL"
	AdaccountcustomaudiencesClaimObjectiveEnumParamVehicle          AdaccountcustomaudiencesClaimObjectiveEnumParam = "VEHICLE"
	AdaccountcustomaudiencesClaimObjectiveEnumParamVehicleOffer     AdaccountcustomaudiencesClaimObjectiveEnumParam = "VEHICLE_OFFER"
)

func (value AdaccountcustomaudiencesClaimObjectiveEnumParam) String() string {
	return string(value)
}
