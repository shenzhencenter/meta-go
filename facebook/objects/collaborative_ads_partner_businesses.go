package objects

type CollaborativeAdsPartnerBusinesses struct {
	CollaborativeAdsPartnerBusinessesInfo *[]Business `json:"collaborative_ads_partner_businesses_info,omitempty"`
	DedicatedPartnerBusinessInfo          *Business   `json:"dedicated_partner_business_info,omitempty"`
	DedicatedPartnersBusinessInfo         *[]Business `json:"dedicated_partners_business_info,omitempty"`
}

var CollaborativeAdsPartnerBusinessesFields = struct {
	CollaborativeAdsPartnerBusinessesInfo string
	DedicatedPartnerBusinessInfo          string
	DedicatedPartnersBusinessInfo         string
}{
	CollaborativeAdsPartnerBusinessesInfo: "collaborative_ads_partner_businesses_info",
	DedicatedPartnerBusinessInfo:          "dedicated_partner_business_info",
	DedicatedPartnersBusinessInfo:         "dedicated_partners_business_info",
}
