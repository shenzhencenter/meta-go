package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type BusinessApplicationRequest struct {
	Application *Application `json:"application,omitempty"`
	ID          *core.ID     `json:"id,omitempty"`
}
