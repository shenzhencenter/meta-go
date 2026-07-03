package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type InsightsResult struct {
	Description           *string          `json:"description,omitempty"`
	DescriptionFromAPIDoc *string          `json:"description_from_api_doc,omitempty"`
	ID                    *core.ID         `json:"id,omitempty"`
	Name                  *string          `json:"name,omitempty"`
	Period                *string          `json:"period,omitempty"`
	Title                 *string          `json:"title,omitempty"`
	Values                *[]InsightsValue `json:"values,omitempty"`
}

var InsightsResultFields = struct {
	Description           string
	DescriptionFromAPIDoc string
	ID                    string
	Name                  string
	Period                string
	Title                 string
	Values                string
}{
	Description:           "description",
	DescriptionFromAPIDoc: "description_from_api_doc",
	ID:                    "id",
	Name:                  "name",
	Period:                "period",
	Title:                 "title",
	Values:                "values",
}
