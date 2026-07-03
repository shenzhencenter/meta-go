package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type CustomConversion struct {
	AccountID                *core.ID                               `json:"account_id,omitempty"`
	AggregationRule          *string                                `json:"aggregation_rule,omitempty"`
	Business                 *Business                              `json:"business,omitempty"`
	CreationTime             *core.Time                             `json:"creation_time,omitempty"`
	CustomEventType          *enums.CustomconversionCustomEventType `json:"custom_event_type,omitempty"`
	DataSources              *[]ExternalEventSource                 `json:"data_sources,omitempty"`
	DefaultConversionValue   *int                                   `json:"default_conversion_value,omitempty"`
	Description              *string                                `json:"description,omitempty"`
	EventSourceType          *string                                `json:"event_source_type,omitempty"`
	FirstFiredTime           *core.Time                             `json:"first_fired_time,omitempty"`
	ID                       *core.ID                               `json:"id,omitempty"`
	IsArchived               *bool                                  `json:"is_archived,omitempty"`
	IsUnavailable            *bool                                  `json:"is_unavailable,omitempty"`
	LastFiredTime            *core.Time                             `json:"last_fired_time,omitempty"`
	Name                     *string                                `json:"name,omitempty"`
	OfflineConversionDataSet *OfflineConversionDataSet              `json:"offline_conversion_data_set,omitempty"`
	Pixel                    *AdsPixel                              `json:"pixel,omitempty"`
	RetentionDays            *uint64                                `json:"retention_days,omitempty"`
	Rule                     *string                                `json:"rule,omitempty"`
}

var CustomConversionFields = struct {
	AccountID                string
	AggregationRule          string
	Business                 string
	CreationTime             string
	CustomEventType          string
	DataSources              string
	DefaultConversionValue   string
	Description              string
	EventSourceType          string
	FirstFiredTime           string
	ID                       string
	IsArchived               string
	IsUnavailable            string
	LastFiredTime            string
	Name                     string
	OfflineConversionDataSet string
	Pixel                    string
	RetentionDays            string
	Rule                     string
}{
	AccountID:                "account_id",
	AggregationRule:          "aggregation_rule",
	Business:                 "business",
	CreationTime:             "creation_time",
	CustomEventType:          "custom_event_type",
	DataSources:              "data_sources",
	DefaultConversionValue:   "default_conversion_value",
	Description:              "description",
	EventSourceType:          "event_source_type",
	FirstFiredTime:           "first_fired_time",
	ID:                       "id",
	IsArchived:               "is_archived",
	IsUnavailable:            "is_unavailable",
	LastFiredTime:            "last_fired_time",
	Name:                     "name",
	OfflineConversionDataSet: "offline_conversion_data_set",
	Pixel:                    "pixel",
	RetentionDays:            "retention_days",
	Rule:                     "rule",
}
