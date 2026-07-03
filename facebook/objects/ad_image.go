package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"time"
)

type AdImage struct {
	AccountID                       *core.ID             `json:"account_id,omitempty"`
	CreatedTime                     *time.Time           `json:"created_time,omitempty"`
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
	UpdatedTime                     *time.Time           `json:"updated_time,omitempty"`
	URL                             *string              `json:"url,omitempty"`
	URLX128                         *string              `json:"url_128,omitempty"`
	Width                           *uint64              `json:"width,omitempty"`
}
