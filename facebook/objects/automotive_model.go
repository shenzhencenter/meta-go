package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AutomotiveModel struct {
	Applinks           *CatalogItemAppLinks                   `json:"applinks,omitempty"`
	AutomotiveModelID  *core.ID                               `json:"automotive_model_id,omitempty"`
	Availability       *string                                `json:"availability,omitempty"`
	BodyStyle          *string                                `json:"body_style,omitempty"`
	Currency           *string                                `json:"currency,omitempty"`
	CustomLabelX0      *string                                `json:"custom_label_0,omitempty"`
	CustomLabelX1      *string                                `json:"custom_label_1,omitempty"`
	CustomLabelX2      *string                                `json:"custom_label_2,omitempty"`
	CustomLabelX3      *string                                `json:"custom_label_3,omitempty"`
	CustomLabelX4      *string                                `json:"custom_label_4,omitempty"`
	CustomNumberX0     *uint64                                `json:"custom_number_0,omitempty"`
	CustomNumberX1     *uint64                                `json:"custom_number_1,omitempty"`
	CustomNumberX2     *uint64                                `json:"custom_number_2,omitempty"`
	CustomNumberX3     *uint64                                `json:"custom_number_3,omitempty"`
	CustomNumberX4     *uint64                                `json:"custom_number_4,omitempty"`
	Description        *string                                `json:"description,omitempty"`
	Drivetrain         *string                                `json:"drivetrain,omitempty"`
	ExteriorColor      *string                                `json:"exterior_color,omitempty"`
	FinanceDescription *string                                `json:"finance_description,omitempty"`
	FinanceType        *string                                `json:"finance_type,omitempty"`
	FuelType           *string                                `json:"fuel_type,omitempty"`
	Generation         *string                                `json:"generation,omitempty"`
	ID                 *core.ID                               `json:"id,omitempty"`
	ImageFetchStatus   *enums.AutomotivemodelImageFetchStatus `json:"image_fetch_status,omitempty"`
	Images             *[]string                              `json:"images,omitempty"`
	InteriorColor      *string                                `json:"interior_color,omitempty"`
	InteriorUpholstery *string                                `json:"interior_upholstery,omitempty"`
	Make               *string                                `json:"make,omitempty"`
	Model              *string                                `json:"model,omitempty"`
	Price              *string                                `json:"price,omitempty"`
	SanitizedImages    *[]string                              `json:"sanitized_images,omitempty"`
	Title              *string                                `json:"title,omitempty"`
	Transmission       *string                                `json:"transmission,omitempty"`
	Trim               *string                                `json:"trim,omitempty"`
	UnitPrice          *map[string]interface{}                `json:"unit_price,omitempty"`
	URL                *string                                `json:"url,omitempty"`
	Visibility         *enums.AutomotivemodelVisibility       `json:"visibility,omitempty"`
	Year               *int                                   `json:"year,omitempty"`
}

var AutomotiveModelFields = struct {
	Applinks           string
	AutomotiveModelID  string
	Availability       string
	BodyStyle          string
	Currency           string
	CustomLabelX0      string
	CustomLabelX1      string
	CustomLabelX2      string
	CustomLabelX3      string
	CustomLabelX4      string
	CustomNumberX0     string
	CustomNumberX1     string
	CustomNumberX2     string
	CustomNumberX3     string
	CustomNumberX4     string
	Description        string
	Drivetrain         string
	ExteriorColor      string
	FinanceDescription string
	FinanceType        string
	FuelType           string
	Generation         string
	ID                 string
	ImageFetchStatus   string
	Images             string
	InteriorColor      string
	InteriorUpholstery string
	Make               string
	Model              string
	Price              string
	SanitizedImages    string
	Title              string
	Transmission       string
	Trim               string
	UnitPrice          string
	URL                string
	Visibility         string
	Year               string
}{
	Applinks:           "applinks",
	AutomotiveModelID:  "automotive_model_id",
	Availability:       "availability",
	BodyStyle:          "body_style",
	Currency:           "currency",
	CustomLabelX0:      "custom_label_0",
	CustomLabelX1:      "custom_label_1",
	CustomLabelX2:      "custom_label_2",
	CustomLabelX3:      "custom_label_3",
	CustomLabelX4:      "custom_label_4",
	CustomNumberX0:     "custom_number_0",
	CustomNumberX1:     "custom_number_1",
	CustomNumberX2:     "custom_number_2",
	CustomNumberX3:     "custom_number_3",
	CustomNumberX4:     "custom_number_4",
	Description:        "description",
	Drivetrain:         "drivetrain",
	ExteriorColor:      "exterior_color",
	FinanceDescription: "finance_description",
	FinanceType:        "finance_type",
	FuelType:           "fuel_type",
	Generation:         "generation",
	ID:                 "id",
	ImageFetchStatus:   "image_fetch_status",
	Images:             "images",
	InteriorColor:      "interior_color",
	InteriorUpholstery: "interior_upholstery",
	Make:               "make",
	Model:              "model",
	Price:              "price",
	SanitizedImages:    "sanitized_images",
	Title:              "title",
	Transmission:       "transmission",
	Trim:               "trim",
	UnitPrice:          "unit_price",
	URL:                "url",
	Visibility:         "visibility",
	Year:               "year",
}
