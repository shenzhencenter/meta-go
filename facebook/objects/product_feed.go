package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ProductFeed struct {
	Country             *string                                `json:"country,omitempty"`
	CreatedTime         *core.Time                             `json:"created_time,omitempty"`
	DefaultCurrency     *string                                `json:"default_currency,omitempty"`
	DeletionEnabled     *bool                                  `json:"deletion_enabled,omitempty"`
	Delimiter           *enums.ProductfeedDelimiterEnum        `json:"delimiter,omitempty"`
	Encoding            *string                                `json:"encoding,omitempty"`
	FileName            *string                                `json:"file_name,omitempty"`
	ID                  *core.ID                               `json:"id,omitempty"`
	IngestionSourceType *enums.ProductfeedIngestionSourceType  `json:"ingestion_source_type,omitempty"`
	ItemSubType         *string                                `json:"item_sub_type,omitempty"`
	LatestUpload        *ProductFeedUpload                     `json:"latest_upload,omitempty"`
	MigratedFromFeedID  *core.ID                               `json:"migrated_from_feed_id,omitempty"`
	Name                *string                                `json:"name,omitempty"`
	OverrideType        *string                                `json:"override_type,omitempty"`
	PrimaryFeeds        *[]string                              `json:"primary_feeds,omitempty"`
	ProductCount        *int                                   `json:"product_count,omitempty"`
	QuotedFieldsMode    *enums.ProductfeedQuotedFieldsModeEnum `json:"quoted_fields_mode,omitempty"`
	Schedule            *ProductFeedSchedule                   `json:"schedule,omitempty"`
	SupplementaryFeeds  *[]string                              `json:"supplementary_feeds,omitempty"`
	UpdateSchedule      *ProductFeedSchedule                   `json:"update_schedule,omitempty"`
}

var ProductFeedFields = struct {
	Country             string
	CreatedTime         string
	DefaultCurrency     string
	DeletionEnabled     string
	Delimiter           string
	Encoding            string
	FileName            string
	ID                  string
	IngestionSourceType string
	ItemSubType         string
	LatestUpload        string
	MigratedFromFeedID  string
	Name                string
	OverrideType        string
	PrimaryFeeds        string
	ProductCount        string
	QuotedFieldsMode    string
	Schedule            string
	SupplementaryFeeds  string
	UpdateSchedule      string
}{
	Country:             "country",
	CreatedTime:         "created_time",
	DefaultCurrency:     "default_currency",
	DeletionEnabled:     "deletion_enabled",
	Delimiter:           "delimiter",
	Encoding:            "encoding",
	FileName:            "file_name",
	ID:                  "id",
	IngestionSourceType: "ingestion_source_type",
	ItemSubType:         "item_sub_type",
	LatestUpload:        "latest_upload",
	MigratedFromFeedID:  "migrated_from_feed_id",
	Name:                "name",
	OverrideType:        "override_type",
	PrimaryFeeds:        "primary_feeds",
	ProductCount:        "product_count",
	QuotedFieldsMode:    "quoted_fields_mode",
	Schedule:            "schedule",
	SupplementaryFeeds:  "supplementary_feeds",
	UpdateSchedule:      "update_schedule",
}
