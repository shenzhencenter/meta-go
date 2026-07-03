package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AsyncRequest struct {
	ID     *core.ID `json:"id,omitempty"`
	Result *string  `json:"result,omitempty"`
	Status *int     `json:"status,omitempty"`
	Type   *int     `json:"type,omitempty"`
}

var AsyncRequestFields = struct {
	ID     string
	Result string
	Status string
	Type   string
}{
	ID:     "id",
	Result: "result",
	Status: "status",
	Type:   "type",
}
