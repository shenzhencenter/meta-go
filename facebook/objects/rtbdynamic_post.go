package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type RTBDynamicPost struct {
	ChildAttachments *[]DynamicPostChildAttachment `json:"child_attachments,omitempty"`
	Created          *time.Time                    `json:"created,omitempty"`
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
