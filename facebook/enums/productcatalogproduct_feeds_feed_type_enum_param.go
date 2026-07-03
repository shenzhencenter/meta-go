package enums

type ProductcatalogproductFeedsFeedTypeEnumParam string

const (
	ProductcatalogproductFeedsFeedTypeEnumParamActivity                 ProductcatalogproductFeedsFeedTypeEnumParam = "ACTIVITY"
	ProductcatalogproductFeedsFeedTypeEnumParamAppAndSoftware           ProductcatalogproductFeedsFeedTypeEnumParam = "APP_AND_SOFTWARE"
	ProductcatalogproductFeedsFeedTypeEnumParamArticleAndPublication    ProductcatalogproductFeedsFeedTypeEnumParam = "ARTICLE_AND_PUBLICATION"
	ProductcatalogproductFeedsFeedTypeEnumParamAutomotiveModel          ProductcatalogproductFeedsFeedTypeEnumParam = "AUTOMOTIVE_MODEL"
	ProductcatalogproductFeedsFeedTypeEnumParamCollection               ProductcatalogproductFeedsFeedTypeEnumParam = "COLLECTION"
	ProductcatalogproductFeedsFeedTypeEnumParamDestination              ProductcatalogproductFeedsFeedTypeEnumParam = "DESTINATION"
	ProductcatalogproductFeedsFeedTypeEnumParamFlight                   ProductcatalogproductFeedsFeedTypeEnumParam = "FLIGHT"
	ProductcatalogproductFeedsFeedTypeEnumParamHomeListing              ProductcatalogproductFeedsFeedTypeEnumParam = "HOME_LISTING"
	ProductcatalogproductFeedsFeedTypeEnumParamHotel                    ProductcatalogproductFeedsFeedTypeEnumParam = "HOTEL"
	ProductcatalogproductFeedsFeedTypeEnumParamHotelRoom                ProductcatalogproductFeedsFeedTypeEnumParam = "HOTEL_ROOM"
	ProductcatalogproductFeedsFeedTypeEnumParamLocalInventory           ProductcatalogproductFeedsFeedTypeEnumParam = "LOCAL_INVENTORY"
	ProductcatalogproductFeedsFeedTypeEnumParamMediaTitle               ProductcatalogproductFeedsFeedTypeEnumParam = "MEDIA_TITLE"
	ProductcatalogproductFeedsFeedTypeEnumParamOffer                    ProductcatalogproductFeedsFeedTypeEnumParam = "OFFER"
	ProductcatalogproductFeedsFeedTypeEnumParamProducts                 ProductcatalogproductFeedsFeedTypeEnumParam = "PRODUCTS"
	ProductcatalogproductFeedsFeedTypeEnumParamProductRatingsAndReviews ProductcatalogproductFeedsFeedTypeEnumParam = "PRODUCT_RATINGS_AND_REVIEWS"
	ProductcatalogproductFeedsFeedTypeEnumParamService                  ProductcatalogproductFeedsFeedTypeEnumParam = "SERVICE"
	ProductcatalogproductFeedsFeedTypeEnumParamTransactableItems        ProductcatalogproductFeedsFeedTypeEnumParam = "TRANSACTABLE_ITEMS"
	ProductcatalogproductFeedsFeedTypeEnumParamVehicles                 ProductcatalogproductFeedsFeedTypeEnumParam = "VEHICLES"
	ProductcatalogproductFeedsFeedTypeEnumParamVehicleOffer             ProductcatalogproductFeedsFeedTypeEnumParam = "VEHICLE_OFFER"
)

func (value ProductcatalogproductFeedsFeedTypeEnumParam) String() string {
	return string(value)
}
