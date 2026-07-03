package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CPASCollaborationRequest struct {
	AdAccountID            *core.ID        `json:"ad_account_id,omitempty"`
	AdAccountName          *string         `json:"ad_account_name,omitempty"`
	Brands                 *[]string       `json:"brands,omitempty"`
	CatalogSegment         *ProductCatalog `json:"catalog_segment,omitempty"`
	ContactEmail           *string         `json:"contact_email,omitempty"`
	ContactFirstName       *string         `json:"contact_first_name,omitempty"`
	ContactLastName        *string         `json:"contact_last_name,omitempty"`
	CreationTime           *core.Time      `json:"creation_time,omitempty"`
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

var CPASCollaborationRequestFields = struct {
	AdAccountID            string
	AdAccountName          string
	Brands                 string
	CatalogSegment         string
	ContactEmail           string
	ContactFirstName       string
	ContactLastName        string
	CreationTime           string
	ID                     string
	PhoneNumber            string
	ReceiverBusiness       string
	RequesterAgencyOrBrand string
	SellerID               string
	SenderBusiness         string
	SenderClientBusiness   string
	ShopURL                string
	Source                 string
	Status                 string
}{
	AdAccountID:            "ad_account_id",
	AdAccountName:          "ad_account_name",
	Brands:                 "brands",
	CatalogSegment:         "catalog_segment",
	ContactEmail:           "contact_email",
	ContactFirstName:       "contact_first_name",
	ContactLastName:        "contact_last_name",
	CreationTime:           "creation_time",
	ID:                     "id",
	PhoneNumber:            "phone_number",
	ReceiverBusiness:       "receiver_business",
	RequesterAgencyOrBrand: "requester_agency_or_brand",
	SellerID:               "seller_id",
	SenderBusiness:         "sender_business",
	SenderClientBusiness:   "sender_client_business",
	ShopURL:                "shop_url",
	Source:                 "source",
	Status:                 "status",
}
