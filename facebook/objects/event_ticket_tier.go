package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type EventTicketTier struct {
	Currency           *string    `json:"currency,omitempty"`
	Description        *string    `json:"description,omitempty"`
	EndSalesTime       *core.Time `json:"end_sales_time,omitempty"`
	EndShowTime        *core.Time `json:"end_show_time,omitempty"`
	FeeSettings        *string    `json:"fee_settings,omitempty"`
	ID                 *core.ID   `json:"id,omitempty"`
	MaximumQuantity    *int       `json:"maximum_quantity,omitempty"`
	Metadata           *string    `json:"metadata,omitempty"`
	MinimumQuantity    *int       `json:"minimum_quantity,omitempty"`
	Name               *string    `json:"name,omitempty"`
	Price              *int       `json:"price,omitempty"`
	Priority           *int       `json:"priority,omitempty"`
	RetailerID         *core.ID   `json:"retailer_id,omitempty"`
	SeatingMapImageURL *string    `json:"seating_map_image_url,omitempty"`
	StartSalesTime     *core.Time `json:"start_sales_time,omitempty"`
	StartShowTime      *core.Time `json:"start_show_time,omitempty"`
	Status             *string    `json:"status,omitempty"`
	TotalQuantity      *int       `json:"total_quantity,omitempty"`
}

var EventTicketTierFields = struct {
	Currency           string
	Description        string
	EndSalesTime       string
	EndShowTime        string
	FeeSettings        string
	ID                 string
	MaximumQuantity    string
	Metadata           string
	MinimumQuantity    string
	Name               string
	Price              string
	Priority           string
	RetailerID         string
	SeatingMapImageURL string
	StartSalesTime     string
	StartShowTime      string
	Status             string
	TotalQuantity      string
}{
	Currency:           "currency",
	Description:        "description",
	EndSalesTime:       "end_sales_time",
	EndShowTime:        "end_show_time",
	FeeSettings:        "fee_settings",
	ID:                 "id",
	MaximumQuantity:    "maximum_quantity",
	Metadata:           "metadata",
	MinimumQuantity:    "minimum_quantity",
	Name:               "name",
	Price:              "price",
	Priority:           "priority",
	RetailerID:         "retailer_id",
	SeatingMapImageURL: "seating_map_image_url",
	StartSalesTime:     "start_sales_time",
	StartShowTime:      "start_show_time",
	Status:             "status",
	TotalQuantity:      "total_quantity",
}
