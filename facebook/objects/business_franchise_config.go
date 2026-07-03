package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessFranchiseConfig struct {
	ActivePartnerCount        *int                `json:"active_partner_count,omitempty"`
	AgencyBusiness            *Business           `json:"agency_business,omitempty"`
	AgencyBusinessAssetGroup  *BusinessAssetGroup `json:"agency_business_asset_group,omitempty"`
	BrandName                 *string             `json:"brand_name,omitempty"`
	Business                  *Business           `json:"business,omitempty"`
	BusinessVertical          *string             `json:"business_vertical,omitempty"`
	ID                        *core.ID            `json:"id,omitempty"`
	PartnerCount              *int                `json:"partner_count,omitempty"`
	PendingAgencyBusiness     *string             `json:"pending_agency_business,omitempty"`
	ProgramCount              *int                `json:"program_count,omitempty"`
	SharedBusinessAssetGroup  *BusinessAssetGroup `json:"shared_business_asset_group,omitempty"`
	SharedCreativeFolderCount *int                `json:"shared_creative_folder_count,omitempty"`
	SharedCustomAudienceCount *int                `json:"shared_custom_audience_count,omitempty"`
}

var BusinessFranchiseConfigFields = struct {
	ActivePartnerCount        string
	AgencyBusiness            string
	AgencyBusinessAssetGroup  string
	BrandName                 string
	Business                  string
	BusinessVertical          string
	ID                        string
	PartnerCount              string
	PendingAgencyBusiness     string
	ProgramCount              string
	SharedBusinessAssetGroup  string
	SharedCreativeFolderCount string
	SharedCustomAudienceCount string
}{
	ActivePartnerCount:        "active_partner_count",
	AgencyBusiness:            "agency_business",
	AgencyBusinessAssetGroup:  "agency_business_asset_group",
	BrandName:                 "brand_name",
	Business:                  "business",
	BusinessVertical:          "business_vertical",
	ID:                        "id",
	PartnerCount:              "partner_count",
	PendingAgencyBusiness:     "pending_agency_business",
	ProgramCount:              "program_count",
	SharedBusinessAssetGroup:  "shared_business_asset_group",
	SharedCreativeFolderCount: "shared_creative_folder_count",
	SharedCustomAudienceCount: "shared_custom_audience_count",
}
