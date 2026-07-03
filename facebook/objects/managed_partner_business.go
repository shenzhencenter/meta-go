package objects

type ManagedPartnerBusiness struct {
	AdAccount            *AdAccount                    `json:"ad_account,omitempty"`
	CatalogSegment       *ProductCatalog               `json:"catalog_segment,omitempty"`
	ExtendedCredit       *ManagedPartnerExtendedCredit `json:"extended_credit,omitempty"`
	Page                 *Page                         `json:"page,omitempty"`
	SellerBusinessInfo   *map[string]interface{}       `json:"seller_business_info,omitempty"`
	SellerBusinessStatus *string                       `json:"seller_business_status,omitempty"`
	Template             *[]map[string]interface{}     `json:"template,omitempty"`
}
