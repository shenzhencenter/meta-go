package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdImage struct {
	AccountID                       *core.ID             `json:"account_id,omitempty"`
	CreatedTime                     *core.Time           `json:"created_time,omitempty"`
	Creatives                       *[]string            `json:"creatives,omitempty"`
	Hash                            *string              `json:"hash,omitempty"`
	Height                          *uint64              `json:"height,omitempty"`
	ID                              *core.ID             `json:"id,omitempty"`
	IsAssociatedCreativesInAdgroups *bool                `json:"is_associated_creatives_in_adgroups,omitempty"`
	Name                            *string              `json:"name,omitempty"`
	OriginalHeight                  *uint64              `json:"original_height,omitempty"`
	OriginalWidth                   *uint64              `json:"original_width,omitempty"`
	OwnerBusiness                   *Business            `json:"owner_business,omitempty"`
	PermalinkURL                    *string              `json:"permalink_url,omitempty"`
	Status                          *enums.AdimageStatus `json:"status,omitempty"`
	UpdatedTime                     *core.Time           `json:"updated_time,omitempty"`
	URL                             *string              `json:"url,omitempty"`
	URLX128                         *string              `json:"url_128,omitempty"`
	Width                           *uint64              `json:"width,omitempty"`
}

var AdImageFields = struct {
	AccountID                       string
	CreatedTime                     string
	Creatives                       string
	Hash                            string
	Height                          string
	ID                              string
	IsAssociatedCreativesInAdgroups string
	Name                            string
	OriginalHeight                  string
	OriginalWidth                   string
	OwnerBusiness                   string
	PermalinkURL                    string
	Status                          string
	UpdatedTime                     string
	URL                             string
	URLX128                         string
	Width                           string
}{
	AccountID:                       "account_id",
	CreatedTime:                     "created_time",
	Creatives:                       "creatives",
	Hash:                            "hash",
	Height:                          "height",
	ID:                              "id",
	IsAssociatedCreativesInAdgroups: "is_associated_creatives_in_adgroups",
	Name:                            "name",
	OriginalHeight:                  "original_height",
	OriginalWidth:                   "original_width",
	OwnerBusiness:                   "owner_business",
	PermalinkURL:                    "permalink_url",
	Status:                          "status",
	UpdatedTime:                     "updated_time",
	URL:                             "url",
	URLX128:                         "url_128",
	Width:                           "width",
}
