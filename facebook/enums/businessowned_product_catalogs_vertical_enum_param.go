package enums

type BusinessownedProductCatalogsVerticalEnumParam string

const (
	BusinessownedProductCatalogsVerticalEnumParamAdoptablePets           BusinessownedProductCatalogsVerticalEnumParam = "adoptable_pets"
	BusinessownedProductCatalogsVerticalEnumParamAppsAndSoftware         BusinessownedProductCatalogsVerticalEnumParam = "apps_and_software"
	BusinessownedProductCatalogsVerticalEnumParamArticlesAndPublications BusinessownedProductCatalogsVerticalEnumParam = "articles_and_publications"
	BusinessownedProductCatalogsVerticalEnumParamCommerce                BusinessownedProductCatalogsVerticalEnumParam = "commerce"
	BusinessownedProductCatalogsVerticalEnumParamDestinations            BusinessownedProductCatalogsVerticalEnumParam = "destinations"
	BusinessownedProductCatalogsVerticalEnumParamFlights                 BusinessownedProductCatalogsVerticalEnumParam = "flights"
	BusinessownedProductCatalogsVerticalEnumParamGeneric                 BusinessownedProductCatalogsVerticalEnumParam = "generic"
	BusinessownedProductCatalogsVerticalEnumParamHomeListings            BusinessownedProductCatalogsVerticalEnumParam = "home_listings"
	BusinessownedProductCatalogsVerticalEnumParamHotels                  BusinessownedProductCatalogsVerticalEnumParam = "hotels"
	BusinessownedProductCatalogsVerticalEnumParamLocalServiceBusinesses  BusinessownedProductCatalogsVerticalEnumParam = "local_service_businesses"
	BusinessownedProductCatalogsVerticalEnumParamMediaTitles             BusinessownedProductCatalogsVerticalEnumParam = "media_titles"
	BusinessownedProductCatalogsVerticalEnumParamOfferItems              BusinessownedProductCatalogsVerticalEnumParam = "offer_items"
	BusinessownedProductCatalogsVerticalEnumParamOfflineCommerce         BusinessownedProductCatalogsVerticalEnumParam = "offline_commerce"
	BusinessownedProductCatalogsVerticalEnumParamServices                BusinessownedProductCatalogsVerticalEnumParam = "services"
	BusinessownedProductCatalogsVerticalEnumParamTransactableItems       BusinessownedProductCatalogsVerticalEnumParam = "transactable_items"
	BusinessownedProductCatalogsVerticalEnumParamVehicles                BusinessownedProductCatalogsVerticalEnumParam = "vehicles"
)

func (value BusinessownedProductCatalogsVerticalEnumParam) String() string {
	return string(value)
}
