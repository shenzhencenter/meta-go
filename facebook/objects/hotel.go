package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type Hotel struct {
	Address           *string                      `json:"address,omitempty"`
	Applinks          *CatalogItemAppLinks         `json:"applinks,omitempty"`
	Brand             *string                      `json:"brand,omitempty"`
	Category          *string                      `json:"category,omitempty"`
	Currency          *string                      `json:"currency,omitempty"`
	CustomLabelX0     *string                      `json:"custom_label_0,omitempty"`
	CustomLabelX1     *string                      `json:"custom_label_1,omitempty"`
	CustomLabelX2     *string                      `json:"custom_label_2,omitempty"`
	CustomLabelX3     *string                      `json:"custom_label_3,omitempty"`
	CustomLabelX4     *string                      `json:"custom_label_4,omitempty"`
	CustomNumberX0    *int                         `json:"custom_number_0,omitempty"`
	CustomNumberX1    *int                         `json:"custom_number_1,omitempty"`
	CustomNumberX2    *int                         `json:"custom_number_2,omitempty"`
	CustomNumberX3    *int                         `json:"custom_number_3,omitempty"`
	CustomNumberX4    *int                         `json:"custom_number_4,omitempty"`
	Description       *string                      `json:"description,omitempty"`
	GuestRatings      *string                      `json:"guest_ratings,omitempty"`
	HotelID           *core.ID                     `json:"hotel_id,omitempty"`
	ID                *core.ID                     `json:"id,omitempty"`
	ImageFetchStatus  *enums.HotelImageFetchStatus `json:"image_fetch_status,omitempty"`
	Images            *[]string                    `json:"images,omitempty"`
	LowestBasePrice   *string                      `json:"lowest_base_price,omitempty"`
	LoyaltyProgram    *string                      `json:"loyalty_program,omitempty"`
	MarginLevel       *uint64                      `json:"margin_level,omitempty"`
	Name              *string                      `json:"name,omitempty"`
	Phone             *string                      `json:"phone,omitempty"`
	ProductPriorityX0 *float64                     `json:"product_priority_0,omitempty"`
	ProductPriorityX1 *float64                     `json:"product_priority_1,omitempty"`
	ProductPriorityX2 *float64                     `json:"product_priority_2,omitempty"`
	ProductPriorityX3 *float64                     `json:"product_priority_3,omitempty"`
	ProductPriorityX4 *float64                     `json:"product_priority_4,omitempty"`
	SalePrice         *string                      `json:"sale_price,omitempty"`
	SanitizedImages   *[]string                    `json:"sanitized_images,omitempty"`
	StarRating        *float64                     `json:"star_rating,omitempty"`
	Tags              *[]string                    `json:"tags,omitempty"`
	UnitPrice         *map[string]interface{}      `json:"unit_price,omitempty"`
	URL               *string                      `json:"url,omitempty"`
	Visibility        *enums.HotelVisibility       `json:"visibility,omitempty"`
}

var HotelFields = struct {
	Address           string
	Applinks          string
	Brand             string
	Category          string
	Currency          string
	CustomLabelX0     string
	CustomLabelX1     string
	CustomLabelX2     string
	CustomLabelX3     string
	CustomLabelX4     string
	CustomNumberX0    string
	CustomNumberX1    string
	CustomNumberX2    string
	CustomNumberX3    string
	CustomNumberX4    string
	Description       string
	GuestRatings      string
	HotelID           string
	ID                string
	ImageFetchStatus  string
	Images            string
	LowestBasePrice   string
	LoyaltyProgram    string
	MarginLevel       string
	Name              string
	Phone             string
	ProductPriorityX0 string
	ProductPriorityX1 string
	ProductPriorityX2 string
	ProductPriorityX3 string
	ProductPriorityX4 string
	SalePrice         string
	SanitizedImages   string
	StarRating        string
	Tags              string
	UnitPrice         string
	URL               string
	Visibility        string
}{
	Address:           "address",
	Applinks:          "applinks",
	Brand:             "brand",
	Category:          "category",
	Currency:          "currency",
	CustomLabelX0:     "custom_label_0",
	CustomLabelX1:     "custom_label_1",
	CustomLabelX2:     "custom_label_2",
	CustomLabelX3:     "custom_label_3",
	CustomLabelX4:     "custom_label_4",
	CustomNumberX0:    "custom_number_0",
	CustomNumberX1:    "custom_number_1",
	CustomNumberX2:    "custom_number_2",
	CustomNumberX3:    "custom_number_3",
	CustomNumberX4:    "custom_number_4",
	Description:       "description",
	GuestRatings:      "guest_ratings",
	HotelID:           "hotel_id",
	ID:                "id",
	ImageFetchStatus:  "image_fetch_status",
	Images:            "images",
	LowestBasePrice:   "lowest_base_price",
	LoyaltyProgram:    "loyalty_program",
	MarginLevel:       "margin_level",
	Name:              "name",
	Phone:             "phone",
	ProductPriorityX0: "product_priority_0",
	ProductPriorityX1: "product_priority_1",
	ProductPriorityX2: "product_priority_2",
	ProductPriorityX3: "product_priority_3",
	ProductPriorityX4: "product_priority_4",
	SalePrice:         "sale_price",
	SanitizedImages:   "sanitized_images",
	StarRating:        "star_rating",
	Tags:              "tags",
	UnitPrice:         "unit_price",
	URL:               "url",
	Visibility:        "visibility",
}
