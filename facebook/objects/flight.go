package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type Flight struct {
	Applinks           *CatalogItemAppLinks          `json:"applinks,omitempty"`
	Currency           *string                       `json:"currency,omitempty"`
	CustomLabelX0      *string                       `json:"custom_label_0,omitempty"`
	CustomLabelX1      *string                       `json:"custom_label_1,omitempty"`
	CustomLabelX2      *string                       `json:"custom_label_2,omitempty"`
	CustomLabelX3      *string                       `json:"custom_label_3,omitempty"`
	CustomLabelX4      *string                       `json:"custom_label_4,omitempty"`
	CustomNumberX0     *uint64                       `json:"custom_number_0,omitempty"`
	CustomNumberX1     *uint64                       `json:"custom_number_1,omitempty"`
	CustomNumberX2     *uint64                       `json:"custom_number_2,omitempty"`
	CustomNumberX3     *uint64                       `json:"custom_number_3,omitempty"`
	CustomNumberX4     *uint64                       `json:"custom_number_4,omitempty"`
	Description        *string                       `json:"description,omitempty"`
	DestinationAirport *string                       `json:"destination_airport,omitempty"`
	DestinationCity    *string                       `json:"destination_city,omitempty"`
	FlightID           *core.ID                      `json:"flight_id,omitempty"`
	ID                 *core.ID                      `json:"id,omitempty"`
	ImageFetchStatus   *enums.FlightImageFetchStatus `json:"image_fetch_status,omitempty"`
	Images             *[]string                     `json:"images,omitempty"`
	OnewayCurrency     *string                       `json:"oneway_currency,omitempty"`
	OnewayPrice        *string                       `json:"oneway_price,omitempty"`
	OriginAirport      *string                       `json:"origin_airport,omitempty"`
	OriginCity         *string                       `json:"origin_city,omitempty"`
	Price              *string                       `json:"price,omitempty"`
	ProductPriorityX0  *float64                      `json:"product_priority_0,omitempty"`
	ProductPriorityX1  *float64                      `json:"product_priority_1,omitempty"`
	ProductPriorityX2  *float64                      `json:"product_priority_2,omitempty"`
	ProductPriorityX3  *float64                      `json:"product_priority_3,omitempty"`
	ProductPriorityX4  *float64                      `json:"product_priority_4,omitempty"`
	SanitizedImages    *[]string                     `json:"sanitized_images,omitempty"`
	Tags               *[]string                     `json:"tags,omitempty"`
	UnitPrice          *map[string]interface{}       `json:"unit_price,omitempty"`
	URL                *string                       `json:"url,omitempty"`
	Visibility         *enums.FlightVisibility       `json:"visibility,omitempty"`
}

var FlightFields = struct {
	Applinks           string
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
	DestinationAirport string
	DestinationCity    string
	FlightID           string
	ID                 string
	ImageFetchStatus   string
	Images             string
	OnewayCurrency     string
	OnewayPrice        string
	OriginAirport      string
	OriginCity         string
	Price              string
	ProductPriorityX0  string
	ProductPriorityX1  string
	ProductPriorityX2  string
	ProductPriorityX3  string
	ProductPriorityX4  string
	SanitizedImages    string
	Tags               string
	UnitPrice          string
	URL                string
	Visibility         string
}{
	Applinks:           "applinks",
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
	DestinationAirport: "destination_airport",
	DestinationCity:    "destination_city",
	FlightID:           "flight_id",
	ID:                 "id",
	ImageFetchStatus:   "image_fetch_status",
	Images:             "images",
	OnewayCurrency:     "oneway_currency",
	OnewayPrice:        "oneway_price",
	OriginAirport:      "origin_airport",
	OriginCity:         "origin_city",
	Price:              "price",
	ProductPriorityX0:  "product_priority_0",
	ProductPriorityX1:  "product_priority_1",
	ProductPriorityX2:  "product_priority_2",
	ProductPriorityX3:  "product_priority_3",
	ProductPriorityX4:  "product_priority_4",
	SanitizedImages:    "sanitized_images",
	Tags:               "tags",
	UnitPrice:          "unit_price",
	URL:                "url",
	Visibility:         "visibility",
}
