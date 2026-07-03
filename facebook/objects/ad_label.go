package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdLabel struct {
	Account     *AdAccount `json:"account,omitempty"`
	CreatedTime *time.Time `json:"created_time,omitempty"`
	ID          *core.ID   `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	UpdatedTime *time.Time `json:"updated_time,omitempty"`
}
