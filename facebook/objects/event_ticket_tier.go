package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type EventTicketTier struct {
	Currency           *string    `json:"currency,omitempty"`
	Description        *string    `json:"description,omitempty"`
	EndSalesTime       *time.Time `json:"end_sales_time,omitempty"`
	EndShowTime        *time.Time `json:"end_show_time,omitempty"`
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
	StartSalesTime     *time.Time `json:"start_sales_time,omitempty"`
	StartShowTime      *time.Time `json:"start_show_time,omitempty"`
	Status             *string    `json:"status,omitempty"`
	TotalQuantity      *int       `json:"total_quantity,omitempty"`
}
