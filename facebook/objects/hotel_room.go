package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var HotelRoomFields = struct {
	Applinks    string
	BasePrice   string
	Currency    string
	Description string
	ID          string
	Images      string
	MarginLevel string
	Name        string
	RoomID      string
	SalePrice   string
	URL         string
}{
	Applinks:    "applinks",
	BasePrice:   "base_price",
	Currency:    "currency",
	Description: "description",
	ID:          "id",
	Images:      "images",
	MarginLevel: "margin_level",
	Name:        "name",
	RoomID:      "room_id",
	SalePrice:   "sale_price",
	URL:         "url",
}
