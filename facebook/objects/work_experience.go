package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WorkExperience struct {
	Description *string                   `json:"description,omitempty"`
	Employer    *Page                     `json:"employer,omitempty"`
	EndDate     *string                   `json:"end_date,omitempty"`
	From        *User                     `json:"from,omitempty"`
	ID          *core.ID                  `json:"id,omitempty"`
	Location    *Page                     `json:"location,omitempty"`
	Position    *Page                     `json:"position,omitempty"`
	Projects    *[]map[string]interface{} `json:"projects,omitempty"`
	StartDate   *string                   `json:"start_date,omitempty"`
	With        *[]User                   `json:"with,omitempty"`
}

var WorkExperienceFields = struct {
	Description string
	Employer    string
	EndDate     string
	From        string
	ID          string
	Location    string
	Position    string
	Projects    string
	StartDate   string
	With        string
}{
	Description: "description",
	Employer:    "employer",
	EndDate:     "end_date",
	From:        "from",
	ID:          "id",
	Location:    "location",
	Position:    "position",
	Projects:    "projects",
	StartDate:   "start_date",
	With:        "with",
}
