package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdLabel struct {
	Account     *AdAccount `json:"account,omitempty"`
	CreatedTime *core.Time `json:"created_time,omitempty"`
	ID          *core.ID   `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	UpdatedTime *core.Time `json:"updated_time,omitempty"`
}

var AdLabelFields = struct {
	Account     string
	CreatedTime string
	ID          string
	Name        string
	UpdatedTime string
}{
	Account:     "account",
	CreatedTime: "created_time",
	ID:          "id",
	Name:        "name",
	UpdatedTime: "updated_time",
}
