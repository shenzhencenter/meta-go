package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
