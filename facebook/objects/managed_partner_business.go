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

var ManagedPartnerBusinessFields = struct {
	AdAccount            string
	CatalogSegment       string
	ExtendedCredit       string
	Page                 string
	SellerBusinessInfo   string
	SellerBusinessStatus string
	Template             string
}{
	AdAccount:            "ad_account",
	CatalogSegment:       "catalog_segment",
	ExtendedCredit:       "extended_credit",
	Page:                 "page",
	SellerBusinessInfo:   "seller_business_info",
	SellerBusinessStatus: "seller_business_status",
	Template:             "template",
}
