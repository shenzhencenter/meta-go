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
