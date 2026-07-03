package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type LocalServiceBusiness struct {
	Address                *map[string]interface{}                     `json:"address,omitempty"`
	Applinks               *CatalogItemAppLinks                        `json:"applinks,omitempty"`
	Availability           *enums.LocalservicebusinessAvailability     `json:"availability,omitempty"`
	Brand                  *string                                     `json:"brand,omitempty"`
	Category               *string                                     `json:"category,omitempty"`
	Condition              *enums.LocalservicebusinessCondition        `json:"condition,omitempty"`
	CuisineType            *string                                     `json:"cuisine_type,omitempty"`
	Currency               *string                                     `json:"currency,omitempty"`
	CustomLabelX0          *string                                     `json:"custom_label_0,omitempty"`
	CustomLabelX1          *string                                     `json:"custom_label_1,omitempty"`
	CustomLabelX2          *string                                     `json:"custom_label_2,omitempty"`
	CustomLabelX3          *string                                     `json:"custom_label_3,omitempty"`
	CustomLabelX4          *string                                     `json:"custom_label_4,omitempty"`
	CustomNumberX0         *uint64                                     `json:"custom_number_0,omitempty"`
	CustomNumberX1         *uint64                                     `json:"custom_number_1,omitempty"`
	CustomNumberX2         *uint64                                     `json:"custom_number_2,omitempty"`
	CustomNumberX3         *uint64                                     `json:"custom_number_3,omitempty"`
	CustomNumberX4         *uint64                                     `json:"custom_number_4,omitempty"`
	Description            *string                                     `json:"description,omitempty"`
	ExpirationDate         *string                                     `json:"expiration_date,omitempty"`
	Gtin                   *string                                     `json:"gtin,omitempty"`
	ID                     *core.ID                                    `json:"id,omitempty"`
	ImageFetchStatus       *enums.LocalservicebusinessImageFetchStatus `json:"image_fetch_status,omitempty"`
	Images                 *[]string                                   `json:"images,omitempty"`
	LocalInfo              *ProductItemLocalInfo                       `json:"local_info,omitempty"`
	LocalServiceBusinessID *core.ID                                    `json:"local_service_business_id,omitempty"`
	MainLocalInfo          *ProductItemLocalInfo                       `json:"main_local_info,omitempty"`
	Phone                  *string                                     `json:"phone,omitempty"`
	Price                  *string                                     `json:"price,omitempty"`
	PriceRange             *string                                     `json:"price_range,omitempty"`
	RetailerCategory       *string                                     `json:"retailer_category,omitempty"`
	SanitizedImages        *[]string                                   `json:"sanitized_images,omitempty"`
	Size                   *string                                     `json:"size,omitempty"`
	Tags                   *[]string                                   `json:"tags,omitempty"`
	Title                  *string                                     `json:"title,omitempty"`
	UnitPrice              *map[string]interface{}                     `json:"unit_price,omitempty"`
	URL                    *string                                     `json:"url,omitempty"`
	VendorID               *core.ID                                    `json:"vendor_id,omitempty"`
	Visibility             *enums.LocalservicebusinessVisibility       `json:"visibility,omitempty"`
}

var LocalServiceBusinessFields = struct {
	Address                string
	Applinks               string
	Availability           string
	Brand                  string
	Category               string
	Condition              string
	CuisineType            string
	Currency               string
	CustomLabelX0          string
	CustomLabelX1          string
	CustomLabelX2          string
	CustomLabelX3          string
	CustomLabelX4          string
	CustomNumberX0         string
	CustomNumberX1         string
	CustomNumberX2         string
	CustomNumberX3         string
	CustomNumberX4         string
	Description            string
	ExpirationDate         string
	Gtin                   string
	ID                     string
	ImageFetchStatus       string
	Images                 string
	LocalInfo              string
	LocalServiceBusinessID string
	MainLocalInfo          string
	Phone                  string
	Price                  string
	PriceRange             string
	RetailerCategory       string
	SanitizedImages        string
	Size                   string
	Tags                   string
	Title                  string
	UnitPrice              string
	URL                    string
	VendorID               string
	Visibility             string
}{
	Address:                "address",
	Applinks:               "applinks",
	Availability:           "availability",
	Brand:                  "brand",
	Category:               "category",
	Condition:              "condition",
	CuisineType:            "cuisine_type",
	Currency:               "currency",
	CustomLabelX0:          "custom_label_0",
	CustomLabelX1:          "custom_label_1",
	CustomLabelX2:          "custom_label_2",
	CustomLabelX3:          "custom_label_3",
	CustomLabelX4:          "custom_label_4",
	CustomNumberX0:         "custom_number_0",
	CustomNumberX1:         "custom_number_1",
	CustomNumberX2:         "custom_number_2",
	CustomNumberX3:         "custom_number_3",
	CustomNumberX4:         "custom_number_4",
	Description:            "description",
	ExpirationDate:         "expiration_date",
	Gtin:                   "gtin",
	ID:                     "id",
	ImageFetchStatus:       "image_fetch_status",
	Images:                 "images",
	LocalInfo:              "local_info",
	LocalServiceBusinessID: "local_service_business_id",
	MainLocalInfo:          "main_local_info",
	Phone:                  "phone",
	Price:                  "price",
	PriceRange:             "price_range",
	RetailerCategory:       "retailer_category",
	SanitizedImages:        "sanitized_images",
	Size:                   "size",
	Tags:                   "tags",
	Title:                  "title",
	UnitPrice:              "unit_price",
	URL:                    "url",
	VendorID:               "vendor_id",
	Visibility:             "visibility",
}
