package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type EducationExperience struct {
	Classes       *[]Experience `json:"classes,omitempty"`
	Concentration *[]Page       `json:"concentration,omitempty"`
	Degree        *Page         `json:"degree,omitempty"`
	ID            *core.ID      `json:"id,omitempty"`
	School        *Page         `json:"school,omitempty"`
	Type          *string       `json:"type,omitempty"`
	With          *[]User       `json:"with,omitempty"`
	Year          *Page         `json:"year,omitempty"`
}

var EducationExperienceFields = struct {
	Classes       string
	Concentration string
	Degree        string
	ID            string
	School        string
	Type          string
	With          string
	Year          string
}{
	Classes:       "classes",
	Concentration: "concentration",
	Degree:        "degree",
	ID:            "id",
	School:        "school",
	Type:          "type",
	With:          "with",
	Year:          "year",
}
