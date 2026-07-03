package enums

type AdaccountproductAudiencesClaimObjectiveEnumParam string

const (
	AdaccountproductAudiencesClaimObjectiveEnumParamAutomotiveModel  AdaccountproductAudiencesClaimObjectiveEnumParam = "AUTOMOTIVE_MODEL"
	AdaccountproductAudiencesClaimObjectiveEnumParamCollaborativeAds AdaccountproductAudiencesClaimObjectiveEnumParam = "COLLABORATIVE_ADS"
	AdaccountproductAudiencesClaimObjectiveEnumParamHomeListing      AdaccountproductAudiencesClaimObjectiveEnumParam = "HOME_LISTING"
	AdaccountproductAudiencesClaimObjectiveEnumParamMediaTitle       AdaccountproductAudiencesClaimObjectiveEnumParam = "MEDIA_TITLE"
	AdaccountproductAudiencesClaimObjectiveEnumParamProduct          AdaccountproductAudiencesClaimObjectiveEnumParam = "PRODUCT"
	AdaccountproductAudiencesClaimObjectiveEnumParamTravel           AdaccountproductAudiencesClaimObjectiveEnumParam = "TRAVEL"
	AdaccountproductAudiencesClaimObjectiveEnumParamVehicle          AdaccountproductAudiencesClaimObjectiveEnumParam = "VEHICLE"
	AdaccountproductAudiencesClaimObjectiveEnumParamVehicleOffer     AdaccountproductAudiencesClaimObjectiveEnumParam = "VEHICLE_OFFER"
)

func (value AdaccountproductAudiencesClaimObjectiveEnumParam) String() string {
	return string(value)
}
