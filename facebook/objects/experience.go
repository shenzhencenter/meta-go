package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Experience struct {
	Description *string                 `json:"description,omitempty"`
	From        *map[string]interface{} `json:"from,omitempty"`
	ID          *core.ID                `json:"id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	With        *[]User                 `json:"with,omitempty"`
}

var ExperienceFields = struct {
	Description string
	From        string
	ID          string
	Name        string
	With        string
}{
	Description: "description",
	From:        "from",
	ID:          "id",
	Name:        "name",
	With:        "with",
}
