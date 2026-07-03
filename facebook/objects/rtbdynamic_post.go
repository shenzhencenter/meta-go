package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type RTBDynamicPost struct {
	ChildAttachments *[]DynamicPostChildAttachment `json:"child_attachments,omitempty"`
	Created          *core.Time                    `json:"created,omitempty"`
	Description      *string                       `json:"description,omitempty"`
	ID               *core.ID                      `json:"id,omitempty"`
	ImageURL         *string                       `json:"image_url,omitempty"`
	Link             *string                       `json:"link,omitempty"`
	Message          *string                       `json:"message,omitempty"`
	OwnerID          *core.ID                      `json:"owner_id,omitempty"`
	PlaceID          *core.ID                      `json:"place_id,omitempty"`
	ProductID        *core.ID                      `json:"product_id,omitempty"`
	Title            *string                       `json:"title,omitempty"`
}

var RTBDynamicPostFields = struct {
	ChildAttachments string
	Created          string
	Description      string
	ID               string
	ImageURL         string
	Link             string
	Message          string
	OwnerID          string
	PlaceID          string
	ProductID        string
	Title            string
}{
	ChildAttachments: "child_attachments",
	Created:          "created",
	Description:      "description",
	ID:               "id",
	ImageURL:         "image_url",
	Link:             "link",
	Message:          "message",
	OwnerID:          "owner_id",
	PlaceID:          "place_id",
	ProductID:        "product_id",
	Title:            "title",
}
