package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BrandedContentShadowIGMediaID struct {
	EligibilityErrors             *[]string      `json:"eligibility_errors,omitempty"`
	HasPermissionForPartnershipAd *bool          `json:"has_permission_for_partnership_ad,omitempty"`
	ID                            *core.ID       `json:"id,omitempty"`
	IsCreatorAllowlisted          *bool          `json:"is_creator_allowlisted,omitempty"`
	LinkedProducts                *[]ProductItem `json:"linked_products,omitempty"`
	OwnerID                       *core.ID       `json:"owner_id,omitempty"`
	Permalink                     *string        `json:"permalink,omitempty"`
	ProductSuggestions            *[]ProductItem `json:"product_suggestions,omitempty"`
	RecommendedCampaignObjectives *[]string      `json:"recommended_campaign_objectives,omitempty"`
}

var BrandedContentShadowIGMediaIDFields = struct {
	EligibilityErrors             string
	HasPermissionForPartnershipAd string
	ID                            string
	IsCreatorAllowlisted          string
	LinkedProducts                string
	OwnerID                       string
	Permalink                     string
	ProductSuggestions            string
	RecommendedCampaignObjectives string
}{
	EligibilityErrors:             "eligibility_errors",
	HasPermissionForPartnershipAd: "has_permission_for_partnership_ad",
	ID:                            "id",
	IsCreatorAllowlisted:          "is_creator_allowlisted",
	LinkedProducts:                "linked_products",
	OwnerID:                       "owner_id",
	Permalink:                     "permalink",
	ProductSuggestions:            "product_suggestions",
	RecommendedCampaignObjectives: "recommended_campaign_objectives",
}
