package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessProject struct {
	Business    *Business               `json:"business,omitempty"`
	CreatedTime *core.Time              `json:"created_time,omitempty"`
	Creator     *map[string]interface{} `json:"creator,omitempty"`
	ID          *core.ID                `json:"id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
}

var BusinessProjectFields = struct {
	Business    string
	CreatedTime string
	Creator     string
	ID          string
	Name        string
}{
	Business:    "business",
	CreatedTime: "created_time",
	Creator:     "creator",
	ID:          "id",
	Name:        "name",
}
