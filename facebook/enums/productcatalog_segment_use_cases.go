package enums

type ProductcatalogSegmentUseCases string

const (
	ProductcatalogSegmentUseCasesAffiliateSellerStorefront             ProductcatalogSegmentUseCases = "AFFILIATE_SELLER_STOREFRONT"
	ProductcatalogSegmentUseCasesAffiliateTaggedOnlyDeprecated         ProductcatalogSegmentUseCases = "AFFILIATE_TAGGED_ONLY_DEPRECATED"
	ProductcatalogSegmentUseCasesCollabAds                             ProductcatalogSegmentUseCases = "COLLAB_ADS"
	ProductcatalogSegmentUseCasesCollabAdsForMarketplacePartner        ProductcatalogSegmentUseCases = "COLLAB_ADS_FOR_MARKETPLACE_PARTNER"
	ProductcatalogSegmentUseCasesCollabAdsSegmentWithoutSegmentSyncing ProductcatalogSegmentUseCases = "COLLAB_ADS_SEGMENT_WITHOUT_SEGMENT_SYNCING"
	ProductcatalogSegmentUseCasesDigitalCirculars                      ProductcatalogSegmentUseCases = "DIGITAL_CIRCULARS"
	ProductcatalogSegmentUseCasesFbLiveShopping                        ProductcatalogSegmentUseCases = "FB_LIVE_SHOPPING"
	ProductcatalogSegmentUseCasesIgShopping                            ProductcatalogSegmentUseCases = "IG_SHOPPING"
	ProductcatalogSegmentUseCasesIgShoppingSuggestedProducts           ProductcatalogSegmentUseCases = "IG_SHOPPING_SUGGESTED_PRODUCTS"
	ProductcatalogSegmentUseCasesMarketplaceShops                      ProductcatalogSegmentUseCases = "MARKETPLACE_SHOPS"
	ProductcatalogSegmentUseCasesTest                                  ProductcatalogSegmentUseCases = "TEST"
)

func (value ProductcatalogSegmentUseCases) String() string {
	return string(value)
}
