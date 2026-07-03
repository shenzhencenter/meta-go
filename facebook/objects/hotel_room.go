package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type HotelRoom struct {
	Applinks    *CatalogItemAppLinks `json:"applinks,omitempty"`
	BasePrice   *string              `json:"base_price,omitempty"`
	Currency    *string              `json:"currency,omitempty"`
	Description *string              `json:"description,omitempty"`
	ID          *core.ID             `json:"id,omitempty"`
	Images      *[]string            `json:"images,omitempty"`
	MarginLevel *string              `json:"margin_level,omitempty"`
	Name        *string              `json:"name,omitempty"`
	RoomID      *core.ID             `json:"room_id,omitempty"`
	SalePrice   *string              `json:"sale_price,omitempty"`
	URL         *string              `json:"url,omitempty"`
}
