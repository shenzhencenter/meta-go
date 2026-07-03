package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type CPASCollaborationRequest struct {
	AdAccountID            *core.ID        `json:"ad_account_id,omitempty"`
	AdAccountName          *string         `json:"ad_account_name,omitempty"`
	Brands                 *[]string       `json:"brands,omitempty"`
	CatalogSegment         *ProductCatalog `json:"catalog_segment,omitempty"`
	ContactEmail           *string         `json:"contact_email,omitempty"`
	ContactFirstName       *string         `json:"contact_first_name,omitempty"`
	ContactLastName        *string         `json:"contact_last_name,omitempty"`
	CreationTime           *time.Time      `json:"creation_time,omitempty"`
	ID                     *core.ID        `json:"id,omitempty"`
	PhoneNumber            *string         `json:"phone_number,omitempty"`
	ReceiverBusiness       *Business       `json:"receiver_business,omitempty"`
	RequesterAgencyOrBrand *string         `json:"requester_agency_or_brand,omitempty"`
	SellerID               *core.ID        `json:"seller_id,omitempty"`
	SenderBusiness         *Business       `json:"sender_business,omitempty"`
	SenderClientBusiness   *Business       `json:"sender_client_business,omitempty"`
	ShopURL                *string         `json:"shop_url,omitempty"`
	Source                 *string         `json:"source,omitempty"`
	Status                 *string         `json:"status,omitempty"`
}
