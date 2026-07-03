package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type Profile struct {
	CanPost     *bool                     `json:"can_post,omitempty"`
	ID          *core.ID                  `json:"id,omitempty"`
	Link        *string                   `json:"link,omitempty"`
	Name        *string                   `json:"name,omitempty"`
	Pic         *string                   `json:"pic,omitempty"`
	PicCrop     *ProfilePictureSource     `json:"pic_crop,omitempty"`
	PicLarge    *string                   `json:"pic_large,omitempty"`
	PicSmall    *string                   `json:"pic_small,omitempty"`
	PicSquare   *string                   `json:"pic_square,omitempty"`
	ProfileType *enums.ProfileProfileType `json:"profile_type,omitempty"`
	Username    *string                   `json:"username,omitempty"`
}
