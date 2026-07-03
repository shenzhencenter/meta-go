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

var OfflineProductItemFields = struct {
	Applinks             string
	Brand                string
	Category             string
	Currency             string
	Description          string
	ID                   string
	ImageFetchStatus     string
	ImageURL             string
	Images               string
	Name                 string
	OfflineProductItemID string
	Price                string
	SanitizedImages      string
	URL                  string
	Visibility           string
}{
	Applinks:             "applinks",
	Brand:                "brand",
	Category:             "category",
	Currency:             "currency",
	Description:          "description",
	ID:                   "id",
	ImageFetchStatus:     "image_fetch_status",
	ImageURL:             "image_url",
	Images:               "images",
	Name:                 "name",
	OfflineProductItemID: "offline_product_item_id",
	Price:                "price",
	SanitizedImages:      "sanitized_images",
	URL:                  "url",
	Visibility:           "visibility",
}
