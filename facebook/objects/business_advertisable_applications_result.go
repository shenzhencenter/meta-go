package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessAdvertisableApplicationsResult struct {
	AreAppEventsUnavailable *bool     `json:"are_app_events_unavailable,omitempty"`
	Business                *Business `json:"business,omitempty"`
	HasInsightPermission    *bool     `json:"has_insight_permission,omitempty"`
	ID                      *core.ID  `json:"id,omitempty"`
	Name                    *string   `json:"name,omitempty"`
	PhotoURL                *string   `json:"photo_url,omitempty"`
}

var BusinessAdvertisableApplicationsResultFields = struct {
	AreAppEventsUnavailable string
	Business                string
	HasInsightPermission    string
	ID                      string
	Name                    string
	PhotoURL                string
}{
	AreAppEventsUnavailable: "are_app_events_unavailable",
	Business:                "business",
	HasInsightPermission:    "has_insight_permission",
	ID:                      "id",
	Name:                    "name",
	PhotoURL:                "photo_url",
}
