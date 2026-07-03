package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
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

var VehicleFields = struct {
	Address                      string
	Applinks                     string
	Availability                 string
	AvailabilityCircleRadius     string
	AvailabilityCircleRadiusUnit string
	BodyStyle                    string
	Condition                    string
	Currency                     string
	CustomLabelX0                string
	CustomLabelX1                string
	CustomLabelX2                string
	CustomLabelX3                string
	CustomLabelX4                string
	CustomNumberX0               string
	CustomNumberX1               string
	CustomNumberX2               string
	CustomNumberX3               string
	CustomNumberX4               string
	DateFirstOnLot               string
	DealerCommunicationChannel   string
	DealerEmail                  string
	DealerID                     string
	DealerName                   string
	DealerPhone                  string
	DealerPrivacyPolicyURL       string
	Description                  string
	Drivetrain                   string
	ExteriorColor                string
	FbPageID                     string
	Features                     string
	FuelType                     string
	ID                           string
	ImageFetchStatus             string
	Images                       string
	InteriorColor                string
	LegalDisclosureImpressumURL  string
	Make                         string
	Mileage                      string
	Model                        string
	PreviousCurrency             string
	PreviousPrice                string
	Price                        string
	ProductPriorityX0            string
	ProductPriorityX1            string
	ProductPriorityX2            string
	ProductPriorityX3            string
	ProductPriorityX4            string
	SaleCurrency                 string
	SalePrice                    string
	SanitizedImages              string
	StateOfVehicle               string
	Tags                         string
	Title                        string
	Transmission                 string
	Trim                         string
	UnitPrice                    string
	URL                          string
	VehicleID                    string
	VehicleRegistrationPlate     string
	VehicleSpecifications        string
	VehicleType                  string
	Vin                          string
	Visibility                   string
	Year                         string
}{
	Address:                      "address",
	Applinks:                     "applinks",
	Availability:                 "availability",
	AvailabilityCircleRadius:     "availability_circle_radius",
	AvailabilityCircleRadiusUnit: "availability_circle_radius_unit",
	BodyStyle:                    "body_style",
	Condition:                    "condition",
	Currency:                     "currency",
	CustomLabelX0:                "custom_label_0",
	CustomLabelX1:                "custom_label_1",
	CustomLabelX2:                "custom_label_2",
	CustomLabelX3:                "custom_label_3",
	CustomLabelX4:                "custom_label_4",
	CustomNumberX0:               "custom_number_0",
	CustomNumberX1:               "custom_number_1",
	CustomNumberX2:               "custom_number_2",
	CustomNumberX3:               "custom_number_3",
	CustomNumberX4:               "custom_number_4",
	DateFirstOnLot:               "date_first_on_lot",
	DealerCommunicationChannel:   "dealer_communication_channel",
	DealerEmail:                  "dealer_email",
	DealerID:                     "dealer_id",
	DealerName:                   "dealer_name",
	DealerPhone:                  "dealer_phone",
	DealerPrivacyPolicyURL:       "dealer_privacy_policy_url",
	Description:                  "description",
	Drivetrain:                   "drivetrain",
	ExteriorColor:                "exterior_color",
	FbPageID:                     "fb_page_id",
	Features:                     "features",
	FuelType:                     "fuel_type",
	ID:                           "id",
	ImageFetchStatus:             "image_fetch_status",
	Images:                       "images",
	InteriorColor:                "interior_color",
	LegalDisclosureImpressumURL:  "legal_disclosure_impressum_url",
	Make:                         "make",
	Mileage:                      "mileage",
	Model:                        "model",
	PreviousCurrency:             "previous_currency",
	PreviousPrice:                "previous_price",
	Price:                        "price",
	ProductPriorityX0:            "product_priority_0",
	ProductPriorityX1:            "product_priority_1",
	ProductPriorityX2:            "product_priority_2",
	ProductPriorityX3:            "product_priority_3",
	ProductPriorityX4:            "product_priority_4",
	SaleCurrency:                 "sale_currency",
	SalePrice:                    "sale_price",
	SanitizedImages:              "sanitized_images",
	StateOfVehicle:               "state_of_vehicle",
	Tags:                         "tags",
	Title:                        "title",
	Transmission:                 "transmission",
	Trim:                         "trim",
	UnitPrice:                    "unit_price",
	URL:                          "url",
	VehicleID:                    "vehicle_id",
	VehicleRegistrationPlate:     "vehicle_registration_plate",
	VehicleSpecifications:        "vehicle_specifications",
	VehicleType:                  "vehicle_type",
	Vin:                          "vin",
	Visibility:                   "visibility",
	Year:                         "year",
}
