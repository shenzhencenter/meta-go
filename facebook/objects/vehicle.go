package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type Vehicle struct {
	Address                      *map[string]interface{}        `json:"address,omitempty"`
	Applinks                     *CatalogItemAppLinks           `json:"applinks,omitempty"`
	Availability                 *string                        `json:"availability,omitempty"`
	AvailabilityCircleRadius     *float64                       `json:"availability_circle_radius,omitempty"`
	AvailabilityCircleRadiusUnit *string                        `json:"availability_circle_radius_unit,omitempty"`
	BodyStyle                    *string                        `json:"body_style,omitempty"`
	Condition                    *string                        `json:"condition,omitempty"`
	Currency                     *string                        `json:"currency,omitempty"`
	CustomLabelX0                *string                        `json:"custom_label_0,omitempty"`
	CustomLabelX1                *string                        `json:"custom_label_1,omitempty"`
	CustomLabelX2                *string                        `json:"custom_label_2,omitempty"`
	CustomLabelX3                *string                        `json:"custom_label_3,omitempty"`
	CustomLabelX4                *string                        `json:"custom_label_4,omitempty"`
	CustomNumberX0               *uint64                        `json:"custom_number_0,omitempty"`
	CustomNumberX1               *uint64                        `json:"custom_number_1,omitempty"`
	CustomNumberX2               *uint64                        `json:"custom_number_2,omitempty"`
	CustomNumberX3               *uint64                        `json:"custom_number_3,omitempty"`
	CustomNumberX4               *uint64                        `json:"custom_number_4,omitempty"`
	DateFirstOnLot               *string                        `json:"date_first_on_lot,omitempty"`
	DealerCommunicationChannel   *string                        `json:"dealer_communication_channel,omitempty"`
	DealerEmail                  *string                        `json:"dealer_email,omitempty"`
	DealerID                     *core.ID                       `json:"dealer_id,omitempty"`
	DealerName                   *string                        `json:"dealer_name,omitempty"`
	DealerPhone                  *string                        `json:"dealer_phone,omitempty"`
	DealerPrivacyPolicyURL       *string                        `json:"dealer_privacy_policy_url,omitempty"`
	Description                  *string                        `json:"description,omitempty"`
	Drivetrain                   *string                        `json:"drivetrain,omitempty"`
	ExteriorColor                *string                        `json:"exterior_color,omitempty"`
	FbPageID                     *Page                          `json:"fb_page_id,omitempty"`
	Features                     *[]map[string]interface{}      `json:"features,omitempty"`
	FuelType                     *string                        `json:"fuel_type,omitempty"`
	ID                           *core.ID                       `json:"id,omitempty"`
	ImageFetchStatus             *enums.VehicleImageFetchStatus `json:"image_fetch_status,omitempty"`
	Images                       *[]string                      `json:"images,omitempty"`
	InteriorColor                *string                        `json:"interior_color,omitempty"`
	LegalDisclosureImpressumURL  *string                        `json:"legal_disclosure_impressum_url,omitempty"`
	Make                         *string                        `json:"make,omitempty"`
	Mileage                      *map[string]interface{}        `json:"mileage,omitempty"`
	Model                        *string                        `json:"model,omitempty"`
	PreviousCurrency             *string                        `json:"previous_currency,omitempty"`
	PreviousPrice                *string                        `json:"previous_price,omitempty"`
	Price                        *string                        `json:"price,omitempty"`
	ProductPriorityX0            *float64                       `json:"product_priority_0,omitempty"`
	ProductPriorityX1            *float64                       `json:"product_priority_1,omitempty"`
	ProductPriorityX2            *float64                       `json:"product_priority_2,omitempty"`
	ProductPriorityX3            *float64                       `json:"product_priority_3,omitempty"`
	ProductPriorityX4            *float64                       `json:"product_priority_4,omitempty"`
	SaleCurrency                 *string                        `json:"sale_currency,omitempty"`
	SalePrice                    *string                        `json:"sale_price,omitempty"`
	SanitizedImages              *[]string                      `json:"sanitized_images,omitempty"`
	StateOfVehicle               *string                        `json:"state_of_vehicle,omitempty"`
	Tags                         *[]string                      `json:"tags,omitempty"`
	Title                        *string                        `json:"title,omitempty"`
	Transmission                 *string                        `json:"transmission,omitempty"`
	Trim                         *string                        `json:"trim,omitempty"`
	UnitPrice                    *map[string]interface{}        `json:"unit_price,omitempty"`
	URL                          *string                        `json:"url,omitempty"`
	VehicleID                    *core.ID                       `json:"vehicle_id,omitempty"`
	VehicleRegistrationPlate     *string                        `json:"vehicle_registration_plate,omitempty"`
	VehicleSpecifications        *[]map[string]interface{}      `json:"vehicle_specifications,omitempty"`
	VehicleType                  *string                        `json:"vehicle_type,omitempty"`
	Vin                          *string                        `json:"vin,omitempty"`
	Visibility                   *enums.VehicleVisibility       `json:"visibility,omitempty"`
	Year                         *uint64                        `json:"year,omitempty"`
}
