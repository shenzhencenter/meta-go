package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountPromotableObjects struct {
	PromotableAppIds  *[]core.ID `json:"promotable_app_ids,omitempty"`
	PromotablePageIds *[]core.ID `json:"promotable_page_ids,omitempty"`
	PromotableUrls    *[]string  `json:"promotable_urls,omitempty"`
}
