package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ProductEventStat struct {
	DateStart                            *string                           `json:"date_start,omitempty"`
	DateStop                             *string                           `json:"date_stop,omitempty"`
	DeviceType                           *enums.ProducteventstatDeviceType `json:"device_type,omitempty"`
	Event                                *enums.ProducteventstatEvent      `json:"event,omitempty"`
	EventSource                          *ExternalEventSource              `json:"event_source,omitempty"`
	TotalContentIdsMatchedOtherCatalogs  *int                              `json:"total_content_ids_matched_other_catalogs,omitempty"`
	TotalMatchedContentIds               *core.ID                          `json:"total_matched_content_ids,omitempty"`
	TotalUnmatchedContentIds             *core.ID                          `json:"total_unmatched_content_ids,omitempty"`
	UniqueContentIdsMatchedOtherCatalogs *int                              `json:"unique_content_ids_matched_other_catalogs,omitempty"`
	UniqueMatchedContentIds              *core.ID                          `json:"unique_matched_content_ids,omitempty"`
	UniqueUnmatchedContentIds            *core.ID                          `json:"unique_unmatched_content_ids,omitempty"`
}
