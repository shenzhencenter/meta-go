package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type TransactableItem struct {
	ActionTitle         *string                                 `json:"action_title,omitempty"`
	Applinks            *CatalogItemAppLinks                    `json:"applinks,omitempty"`
	Currency            *string                                 `json:"currency,omitempty"`
	Description         *string                                 `json:"description,omitempty"`
	DurationTime        *uint64                                 `json:"duration_time,omitempty"`
	DurationType        *string                                 `json:"duration_type,omitempty"`
	ID                  *core.ID                                `json:"id,omitempty"`
	ImageFetchStatus    *enums.TransactableitemImageFetchStatus `json:"image_fetch_status,omitempty"`
	Images              *[]string                               `json:"images,omitempty"`
	OrderIndex          *uint64                                 `json:"order_index,omitempty"`
	Price               *string                                 `json:"price,omitempty"`
	PriceType           *string                                 `json:"price_type,omitempty"`
	SanitizedImages     *[]string                               `json:"sanitized_images,omitempty"`
	SessionType         *string                                 `json:"session_type,omitempty"`
	TimePaddingAfterEnd *uint64                                 `json:"time_padding_after_end,omitempty"`
	Title               *string                                 `json:"title,omitempty"`
	TransactableItemID  *core.ID                                `json:"transactable_item_id,omitempty"`
	URL                 *string                                 `json:"url,omitempty"`
	Visibility          *enums.TransactableitemVisibility       `json:"visibility,omitempty"`
}
