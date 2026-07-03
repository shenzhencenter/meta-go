package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
