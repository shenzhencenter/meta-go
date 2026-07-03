package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type HomeListing struct {
	AcType                    *string                            `json:"ac_type,omitempty"`
	AdditionalFeesDescription *string                            `json:"additional_fees_description,omitempty"`
	Address                   *map[string]interface{}            `json:"address,omitempty"`
	AgentCompany              *string                            `json:"agent_company,omitempty"`
	AgentEmail                *string                            `json:"agent_email,omitempty"`
	AgentFbPageID             *Page                              `json:"agent_fb_page_id,omitempty"`
	AgentName                 *string                            `json:"agent_name,omitempty"`
	AgentPhone                *string                            `json:"agent_phone,omitempty"`
	Applinks                  *CatalogItemAppLinks               `json:"applinks,omitempty"`
	AreaSize                  *uint64                            `json:"area_size,omitempty"`
	AreaUnit                  *string                            `json:"area_unit,omitempty"`
	Availability              *string                            `json:"availability,omitempty"`
	CoX2EmissionRatingEu      *map[string]interface{}            `json:"co_2_emission_rating_eu,omitempty"`
	Currency                  *string                            `json:"currency,omitempty"`
	CustomLabelX0             *string                            `json:"custom_label_0,omitempty"`
	CustomLabelX1             *string                            `json:"custom_label_1,omitempty"`
	CustomLabelX2             *string                            `json:"custom_label_2,omitempty"`
	CustomLabelX3             *string                            `json:"custom_label_3,omitempty"`
	CustomLabelX4             *string                            `json:"custom_label_4,omitempty"`
	CustomNumberX0            *uint64                            `json:"custom_number_0,omitempty"`
	CustomNumberX1            *uint64                            `json:"custom_number_1,omitempty"`
	CustomNumberX2            *uint64                            `json:"custom_number_2,omitempty"`
	CustomNumberX3            *uint64                            `json:"custom_number_3,omitempty"`
	CustomNumberX4            *uint64                            `json:"custom_number_4,omitempty"`
	DaysOnMarket              *uint64                            `json:"days_on_market,omitempty"`
	Description               *string                            `json:"description,omitempty"`
	EnergyRatingEu            *map[string]interface{}            `json:"energy_rating_eu,omitempty"`
	FurnishType               *string                            `json:"furnish_type,omitempty"`
	GroupID                   *core.ID                           `json:"group_id,omitempty"`
	HeatingType               *string                            `json:"heating_type,omitempty"`
	HomeListingID             *core.ID                           `json:"home_listing_id,omitempty"`
	ID                        *core.ID                           `json:"id,omitempty"`
	ImageFetchStatus          *enums.HomelistingImageFetchStatus `json:"image_fetch_status,omitempty"`
	Images                    *[]string                          `json:"images,omitempty"`
	LaundryType               *string                            `json:"laundry_type,omitempty"`
	ListingType               *string                            `json:"listing_type,omitempty"`
	MaxCurrency               *string                            `json:"max_currency,omitempty"`
	MaxPrice                  *string                            `json:"max_price,omitempty"`
	MinCurrency               *string                            `json:"min_currency,omitempty"`
	MinPrice                  *string                            `json:"min_price,omitempty"`
	Name                      *string                            `json:"name,omitempty"`
	NumBaths                  *float64                           `json:"num_baths,omitempty"`
	NumBeds                   *float64                           `json:"num_beds,omitempty"`
	NumRooms                  *float64                           `json:"num_rooms,omitempty"`
	NumUnits                  *uint64                            `json:"num_units,omitempty"`
	ParkingType               *string                            `json:"parking_type,omitempty"`
	PartnerVerification       *string                            `json:"partner_verification,omitempty"`
	PetPolicy                 *string                            `json:"pet_policy,omitempty"`
	Price                     *string                            `json:"price,omitempty"`
	PropertyType              *string                            `json:"property_type,omitempty"`
	SanitizedImages           *[]string                          `json:"sanitized_images,omitempty"`
	SecuritydepositCurrency   *string                            `json:"securitydeposit_currency,omitempty"`
	SecuritydepositPrice      *string                            `json:"securitydeposit_price,omitempty"`
	Tags                      *[]string                          `json:"tags,omitempty"`
	UnitPrice                 *map[string]interface{}            `json:"unit_price,omitempty"`
	URL                       *string                            `json:"url,omitempty"`
	Visibility                *enums.HomelistingVisibility       `json:"visibility,omitempty"`
	YearBuilt                 *int                               `json:"year_built,omitempty"`
}
