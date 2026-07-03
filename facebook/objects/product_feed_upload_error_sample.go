package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductFeedUploadErrorSample struct {
	ID         *core.ID `json:"id,omitempty"`
	RetailerID *core.ID `json:"retailer_id,omitempty"`
	RowNumber  *int     `json:"row_number,omitempty"`
}

var ProductFeedUploadErrorSampleFields = struct {
	ID         string
	RetailerID string
	RowNumber  string
}{
	ID:         "id",
	RetailerID: "retailer_id",
	RowNumber:  "row_number",
}
