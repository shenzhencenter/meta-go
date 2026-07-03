package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type OfflineProductItem struct {
	Applinks             *CatalogItemAppLinks                      `json:"applinks,omitempty"`
	Brand                *string                                   `json:"brand,omitempty"`
	Category             *string                                   `json:"category,omitempty"`
	Currency             *string                                   `json:"currency,omitempty"`
	Description          *string                                   `json:"description,omitempty"`
	ID                   *core.ID                                  `json:"id,omitempty"`
	ImageFetchStatus     *enums.OfflineproductitemImageFetchStatus `json:"image_fetch_status,omitempty"`
	ImageURL             *string                                   `json:"image_url,omitempty"`
	Images               *[]string                                 `json:"images,omitempty"`
	Name                 *string                                   `json:"name,omitempty"`
	OfflineProductItemID *core.ID                                  `json:"offline_product_item_id,omitempty"`
	Price                *string                                   `json:"price,omitempty"`
	SanitizedImages      *[]string                                 `json:"sanitized_images,omitempty"`
	URL                  *string                                   `json:"url,omitempty"`
	Visibility           *enums.OfflineproductitemVisibility       `json:"visibility,omitempty"`
}
