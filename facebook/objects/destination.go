package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type Destination struct {
	Address          *string                            `json:"address,omitempty"`
	Applinks         *CatalogItemAppLinks               `json:"applinks,omitempty"`
	Currency         *string                            `json:"currency,omitempty"`
	CustomLabelX0    *string                            `json:"custom_label_0,omitempty"`
	CustomLabelX1    *string                            `json:"custom_label_1,omitempty"`
	CustomLabelX2    *string                            `json:"custom_label_2,omitempty"`
	CustomLabelX3    *string                            `json:"custom_label_3,omitempty"`
	CustomLabelX4    *string                            `json:"custom_label_4,omitempty"`
	CustomNumberX0   *uint64                            `json:"custom_number_0,omitempty"`
	CustomNumberX1   *uint64                            `json:"custom_number_1,omitempty"`
	CustomNumberX2   *uint64                            `json:"custom_number_2,omitempty"`
	CustomNumberX3   *uint64                            `json:"custom_number_3,omitempty"`
	CustomNumberX4   *uint64                            `json:"custom_number_4,omitempty"`
	Description      *string                            `json:"description,omitempty"`
	DestinationID    *core.ID                           `json:"destination_id,omitempty"`
	ID               *core.ID                           `json:"id,omitempty"`
	ImageFetchStatus *enums.DestinationImageFetchStatus `json:"image_fetch_status,omitempty"`
	Images           *[]string                          `json:"images,omitempty"`
	Name             *string                            `json:"name,omitempty"`
	Price            *string                            `json:"price,omitempty"`
	PriceChange      *string                            `json:"price_change,omitempty"`
	SanitizedImages  *[]string                          `json:"sanitized_images,omitempty"`
	Tags             *[]string                          `json:"tags,omitempty"`
	Types            *[]string                          `json:"types,omitempty"`
	UnitPrice        *map[string]interface{}            `json:"unit_price,omitempty"`
	URL              *string                            `json:"url,omitempty"`
	Visibility       *enums.DestinationVisibility       `json:"visibility,omitempty"`
}
