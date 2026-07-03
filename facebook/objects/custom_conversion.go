package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"time"
)

type CustomConversion struct {
	AccountID                *core.ID                               `json:"account_id,omitempty"`
	AggregationRule          *string                                `json:"aggregation_rule,omitempty"`
	Business                 *Business                              `json:"business,omitempty"`
	CreationTime             *time.Time                             `json:"creation_time,omitempty"`
	CustomEventType          *enums.CustomconversionCustomEventType `json:"custom_event_type,omitempty"`
	DataSources              *[]ExternalEventSource                 `json:"data_sources,omitempty"`
	DefaultConversionValue   *int                                   `json:"default_conversion_value,omitempty"`
	Description              *string                                `json:"description,omitempty"`
	EventSourceType          *string                                `json:"event_source_type,omitempty"`
	FirstFiredTime           *time.Time                             `json:"first_fired_time,omitempty"`
	ID                       *core.ID                               `json:"id,omitempty"`
	IsArchived               *bool                                  `json:"is_archived,omitempty"`
	IsUnavailable            *bool                                  `json:"is_unavailable,omitempty"`
	LastFiredTime            *time.Time                             `json:"last_fired_time,omitempty"`
	Name                     *string                                `json:"name,omitempty"`
	OfflineConversionDataSet *OfflineConversionDataSet              `json:"offline_conversion_data_set,omitempty"`
	Pixel                    *AdsPixel                              `json:"pixel,omitempty"`
	RetentionDays            *uint64                                `json:"retention_days,omitempty"`
	Rule                     *string                                `json:"rule,omitempty"`
}
