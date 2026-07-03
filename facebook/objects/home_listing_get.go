package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type HomeListingGet struct {
	AcType                              *enums.HomelistinggetAcType                                `json:"ac_type,omitempty"`
	AdditionalFeesDescription           *string                                                    `json:"additional_fees_description,omitempty"`
	Address                             *map[string]interface{}                                    `json:"address,omitempty"`
	AgentCompany                        *string                                                    `json:"agent_company,omitempty"`
	AgentEmail                          *string                                                    `json:"agent_email,omitempty"`
	AgentFbPageID                       *map[string]interface{}                                    `json:"agent_fb_page_id,omitempty"`
	AgentName                           *string                                                    `json:"agent_name,omitempty"`
	AgentPhone                          *string                                                    `json:"agent_phone,omitempty"`
	Applinks                            *map[string]interface{}                                    `json:"applinks,omitempty"`
	AreaSize                            *int                                                       `json:"area_size,omitempty"`
	AreaUnit                            *enums.HomelistinggetAreaUnit                              `json:"area_unit,omitempty"`
	Availability                        *enums.HomelistinggetAvailability                          `json:"availability,omitempty"`
	Capabilities                        *[]enums.HomelistinggetCapabilities                        `json:"capabilities,omitempty"`
	CapabilitiesDisabledByUser          *[]string                                                  `json:"capabilities_disabled_by_user,omitempty"`
	CapabilitiesRenderingCatalogManager *[]enums.HomelistinggetCapabilitiesRenderingCatalogManager `json:"capabilities_rendering_catalog_manager,omitempty"`
	CapabilityToRejectionReason         *[]map[string]interface{}                                  `json:"capability_to_rejection_reason,omitempty"`
	CatalogItemOverrides                *map[string]interface{}                                    `json:"catalog_item_overrides,omitempty"`
	ChannelsToIntegrityStatus           *map[string]interface{}                                    `json:"channels_to_integrity_status,omitempty"`
	CoX2EmissionRatingEu                *map[string]interface{}                                    `json:"co_2_emission_rating_eu,omitempty"`
	Currency                            *string                                                    `json:"currency,omitempty"`
	CustomLabelX0                       *string                                                    `json:"custom_label_0,omitempty"`
	CustomLabelX1                       *string                                                    `json:"custom_label_1,omitempty"`
	CustomLabelX2                       *string                                                    `json:"custom_label_2,omitempty"`
	CustomLabelX3                       *string                                                    `json:"custom_label_3,omitempty"`
	CustomLabelX4                       *string                                                    `json:"custom_label_4,omitempty"`
	CustomNumberX0                      *int                                                       `json:"custom_number_0,omitempty"`
	CustomNumberX1                      *int                                                       `json:"custom_number_1,omitempty"`
	CustomNumberX2                      *int                                                       `json:"custom_number_2,omitempty"`
	CustomNumberX3                      *int                                                       `json:"custom_number_3,omitempty"`
	CustomNumberX4                      *int                                                       `json:"custom_number_4,omitempty"`
	DaysOnMarket                        *int                                                       `json:"days_on_market,omitempty"`
	Description                         *string                                                    `json:"description,omitempty"`
	EnabledCapabilityToReviewStatus     *[]map[string]interface{}                                  `json:"enabled_capability_to_review_status,omitempty"`
	EnergyRatingEu                      *map[string]interface{}                                    `json:"energy_rating_eu,omitempty"`
	FurnishType                         *enums.HomelistinggetFurnishType                           `json:"furnish_type,omitempty"`
	GroupID                             *core.ID                                                   `json:"group_id,omitempty"`
	HeatingType                         *enums.HomelistinggetHeatingType                           `json:"heating_type,omitempty"`
	HomeListingID                       *core.ID                                                   `json:"home_listing_id,omitempty"`
	ID                                  *core.ID                                                   `json:"id,omitempty"`
	ImageFetchStatus                    *enums.HomelistinggetImageFetchStatus                      `json:"image_fetch_status,omitempty"`
	Images                              *[]string                                                  `json:"images,omitempty"`
	IsBlackholed                        *bool                                                      `json:"is_blackholed,omitempty"`
	LaundryType                         *enums.HomelistinggetLaundryType                           `json:"laundry_type,omitempty"`
	ListingType                         *enums.HomelistinggetListingType                           `json:"listing_type,omitempty"`
	MaxCurrency                         *string                                                    `json:"max_currency,omitempty"`
	MaxPrice                            *string                                                    `json:"max_price,omitempty"`
	MinCurrency                         *string                                                    `json:"min_currency,omitempty"`
	MinPrice                            *string                                                    `json:"min_price,omitempty"`
	Name                                *string                                                    `json:"name,omitempty"`
	NumBaths                            *float64                                                   `json:"num_baths,omitempty"`
	NumBeds                             *float64                                                   `json:"num_beds,omitempty"`
	NumRooms                            *float64                                                   `json:"num_rooms,omitempty"`
	NumUnits                            *int                                                       `json:"num_units,omitempty"`
	OverrideDetails                     *map[string]interface{}                                    `json:"override_details,omitempty"`
	ParkingType                         *enums.HomelistinggetParkingType                           `json:"parking_type,omitempty"`
	PartnerVerification                 *enums.HomelistinggetPartnerVerification                   `json:"partner_verification,omitempty"`
	PetPolicy                           *string                                                    `json:"pet_policy,omitempty"`
	Price                               *string                                                    `json:"price,omitempty"`
	ProductFeed                         *map[string]interface{}                                    `json:"product_feed,omitempty"`
	PropertyType                        *enums.HomelistinggetPropertyType                          `json:"property_type,omitempty"`
	SanitizedImages                     *[]string                                                  `json:"sanitized_images,omitempty"`
	SanitizedPreviews                   *[]string                                                  `json:"sanitized_previews,omitempty"`
	SecuritydepositCurrency             *string                                                    `json:"securitydeposit_currency,omitempty"`
	SecuritydepositPrice                *string                                                    `json:"securitydeposit_price,omitempty"`
	Tags                                *[]string                                                  `json:"tags,omitempty"`
	UnitPrice                           *map[string]interface{}                                    `json:"unit_price,omitempty"`
	URL                                 *string                                                    `json:"url,omitempty"`
	URLShimmed                          *string                                                    `json:"url_shimmed,omitempty"`
	ValidationErrors                    *map[string]interface{}                                    `json:"validation_errors,omitempty"`
	VideosMetadata                      *map[string]interface{}                                    `json:"videos_metadata,omitempty"`
	Visibility                          *enums.HomelistinggetVisibility                            `json:"visibility,omitempty"`
	YearBuilt                           *int                                                       `json:"year_built,omitempty"`
}

var HomeListingGetFields = struct {
	AcType                              string
	AdditionalFeesDescription           string
	Address                             string
	AgentCompany                        string
	AgentEmail                          string
	AgentFbPageID                       string
	AgentName                           string
	AgentPhone                          string
	Applinks                            string
	AreaSize                            string
	AreaUnit                            string
	Availability                        string
	Capabilities                        string
	CapabilitiesDisabledByUser          string
	CapabilitiesRenderingCatalogManager string
	CapabilityToRejectionReason         string
	CatalogItemOverrides                string
	ChannelsToIntegrityStatus           string
	CoX2EmissionRatingEu                string
	Currency                            string
	CustomLabelX0                       string
	CustomLabelX1                       string
	CustomLabelX2                       string
	CustomLabelX3                       string
	CustomLabelX4                       string
	CustomNumberX0                      string
	CustomNumberX1                      string
	CustomNumberX2                      string
	CustomNumberX3                      string
	CustomNumberX4                      string
	DaysOnMarket                        string
	Description                         string
	EnabledCapabilityToReviewStatus     string
	EnergyRatingEu                      string
	FurnishType                         string
	GroupID                             string
	HeatingType                         string
	HomeListingID                       string
	ID                                  string
	ImageFetchStatus                    string
	Images                              string
	IsBlackholed                        string
	LaundryType                         string
	ListingType                         string
	MaxCurrency                         string
	MaxPrice                            string
	MinCurrency                         string
	MinPrice                            string
	Name                                string
	NumBaths                            string
	NumBeds                             string
	NumRooms                            string
	NumUnits                            string
	OverrideDetails                     string
	ParkingType                         string
	PartnerVerification                 string
	PetPolicy                           string
	Price                               string
	ProductFeed                         string
	PropertyType                        string
	SanitizedImages                     string
	SanitizedPreviews                   string
	SecuritydepositCurrency             string
	SecuritydepositPrice                string
	Tags                                string
	UnitPrice                           string
	URL                                 string
	URLShimmed                          string
	ValidationErrors                    string
	VideosMetadata                      string
	Visibility                          string
	YearBuilt                           string
}{
	AcType:                              "ac_type",
	AdditionalFeesDescription:           "additional_fees_description",
	Address:                             "address",
	AgentCompany:                        "agent_company",
	AgentEmail:                          "agent_email",
	AgentFbPageID:                       "agent_fb_page_id",
	AgentName:                           "agent_name",
	AgentPhone:                          "agent_phone",
	Applinks:                            "applinks",
	AreaSize:                            "area_size",
	AreaUnit:                            "area_unit",
	Availability:                        "availability",
	Capabilities:                        "capabilities",
	CapabilitiesDisabledByUser:          "capabilities_disabled_by_user",
	CapabilitiesRenderingCatalogManager: "capabilities_rendering_catalog_manager",
	CapabilityToRejectionReason:         "capability_to_rejection_reason",
	CatalogItemOverrides:                "catalog_item_overrides",
	ChannelsToIntegrityStatus:           "channels_to_integrity_status",
	CoX2EmissionRatingEu:                "co_2_emission_rating_eu",
	Currency:                            "currency",
	CustomLabelX0:                       "custom_label_0",
	CustomLabelX1:                       "custom_label_1",
	CustomLabelX2:                       "custom_label_2",
	CustomLabelX3:                       "custom_label_3",
	CustomLabelX4:                       "custom_label_4",
	CustomNumberX0:                      "custom_number_0",
	CustomNumberX1:                      "custom_number_1",
	CustomNumberX2:                      "custom_number_2",
	CustomNumberX3:                      "custom_number_3",
	CustomNumberX4:                      "custom_number_4",
	DaysOnMarket:                        "days_on_market",
	Description:                         "description",
	EnabledCapabilityToReviewStatus:     "enabled_capability_to_review_status",
	EnergyRatingEu:                      "energy_rating_eu",
	FurnishType:                         "furnish_type",
	GroupID:                             "group_id",
	HeatingType:                         "heating_type",
	HomeListingID:                       "home_listing_id",
	ID:                                  "id",
	ImageFetchStatus:                    "image_fetch_status",
	Images:                              "images",
	IsBlackholed:                        "is_blackholed",
	LaundryType:                         "laundry_type",
	ListingType:                         "listing_type",
	MaxCurrency:                         "max_currency",
	MaxPrice:                            "max_price",
	MinCurrency:                         "min_currency",
	MinPrice:                            "min_price",
	Name:                                "name",
	NumBaths:                            "num_baths",
	NumBeds:                             "num_beds",
	NumRooms:                            "num_rooms",
	NumUnits:                            "num_units",
	OverrideDetails:                     "override_details",
	ParkingType:                         "parking_type",
	PartnerVerification:                 "partner_verification",
	PetPolicy:                           "pet_policy",
	Price:                               "price",
	ProductFeed:                         "product_feed",
	PropertyType:                        "property_type",
	SanitizedImages:                     "sanitized_images",
	SanitizedPreviews:                   "sanitized_previews",
	SecuritydepositCurrency:             "securitydeposit_currency",
	SecuritydepositPrice:                "securitydeposit_price",
	Tags:                                "tags",
	UnitPrice:                           "unit_price",
	URL:                                 "url",
	URLShimmed:                          "url_shimmed",
	ValidationErrors:                    "validation_errors",
	VideosMetadata:                      "videos_metadata",
	Visibility:                          "visibility",
	YearBuilt:                           "year_built",
}
