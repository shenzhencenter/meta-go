package enums

type CustomaudienceClaimObjective string

const (
	CustomaudienceClaimObjectiveAutomotiveModel  CustomaudienceClaimObjective = "AUTOMOTIVE_MODEL"
	CustomaudienceClaimObjectiveCollaborativeAds CustomaudienceClaimObjective = "COLLABORATIVE_ADS"
	CustomaudienceClaimObjectiveHomeListing      CustomaudienceClaimObjective = "HOME_LISTING"
	CustomaudienceClaimObjectiveMediaTitle       CustomaudienceClaimObjective = "MEDIA_TITLE"
	CustomaudienceClaimObjectiveProduct          CustomaudienceClaimObjective = "PRODUCT"
	CustomaudienceClaimObjectiveTravel           CustomaudienceClaimObjective = "TRAVEL"
	CustomaudienceClaimObjectiveVehicle          CustomaudienceClaimObjective = "VEHICLE"
	CustomaudienceClaimObjectiveVehicleOffer     CustomaudienceClaimObjective = "VEHICLE_OFFER"
)

func (value CustomaudienceClaimObjective) String() string {
	return string(value)
}
