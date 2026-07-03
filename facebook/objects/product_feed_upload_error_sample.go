package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ProductFeedUploadErrorSample struct {
	ID         *core.ID `json:"id,omitempty"`
	RetailerID *core.ID `json:"retailer_id,omitempty"`
	RowNumber  *int     `json:"row_number,omitempty"`
}
