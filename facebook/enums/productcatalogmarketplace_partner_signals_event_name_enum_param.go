package enums

type ProductcatalogmarketplacePartnerSignalsEventNameEnumParam string

const (
	ProductcatalogmarketplacePartnerSignalsEventNameEnumParamAddToCart        ProductcatalogmarketplacePartnerSignalsEventNameEnumParam = "ADD_TO_CART"
	ProductcatalogmarketplacePartnerSignalsEventNameEnumParamOfferSubmitted   ProductcatalogmarketplacePartnerSignalsEventNameEnumParam = "OFFER_SUBMITTED"
	ProductcatalogmarketplacePartnerSignalsEventNameEnumParamPurchase         ProductcatalogmarketplacePartnerSignalsEventNameEnumParam = "PURCHASE"
	ProductcatalogmarketplacePartnerSignalsEventNameEnumParamPurchaseViaOffer ProductcatalogmarketplacePartnerSignalsEventNameEnumParam = "PURCHASE_VIA_OFFER"
	ProductcatalogmarketplacePartnerSignalsEventNameEnumParamTest             ProductcatalogmarketplacePartnerSignalsEventNameEnumParam = "TEST"
	ProductcatalogmarketplacePartnerSignalsEventNameEnumParamViewItem         ProductcatalogmarketplacePartnerSignalsEventNameEnumParam = "VIEW_ITEM"
)

func (value ProductcatalogmarketplacePartnerSignalsEventNameEnumParam) String() string {
	return string(value)
}
