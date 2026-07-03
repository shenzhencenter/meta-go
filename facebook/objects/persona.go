package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Persona struct {
	ID                *core.ID `json:"id,omitempty"`
	Name              *string  `json:"name,omitempty"`
	ProfilePictureURL *string  `json:"profile_picture_url,omitempty"`
}
