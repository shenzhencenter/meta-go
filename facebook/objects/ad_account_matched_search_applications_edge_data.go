package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountMatchedSearchApplicationsEdgeData struct {
	AppID                   *core.ID `json:"app_id,omitempty"`
	AreAppEventsUnavailable *bool    `json:"are_app_events_unavailable,omitempty"`
	IconURL                 *string  `json:"icon_url,omitempty"`
	Name                    *string  `json:"name,omitempty"`
	SearchSourceStore       *string  `json:"search_source_store,omitempty"`
	Store                   *string  `json:"store,omitempty"`
	UniqueID                *core.ID `json:"unique_id,omitempty"`
	URL                     *string  `json:"url,omitempty"`
}

var AdAccountMatchedSearchApplicationsEdgeDataFields = struct {
	AppID                   string
	AreAppEventsUnavailable string
	IconURL                 string
	Name                    string
	SearchSourceStore       string
	Store                   string
	UniqueID                string
	URL                     string
}{
	AppID:                   "app_id",
	AreAppEventsUnavailable: "are_app_events_unavailable",
	IconURL:                 "icon_url",
	Name:                    "name",
	SearchSourceStore:       "search_source_store",
	Store:                   "store",
	UniqueID:                "unique_id",
	URL:                     "url",
}
