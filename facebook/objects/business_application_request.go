package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessApplicationRequest struct {
	Application *Application `json:"application,omitempty"`
	ID          *core.ID     `json:"id,omitempty"`
}

var BusinessApplicationRequestFields = struct {
	Application string
	ID          string
}{
	Application: "application",
	ID:          "id",
}
