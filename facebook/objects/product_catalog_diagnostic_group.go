package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ProductCatalogDiagnosticGroup struct {
	AffectedChannels         *[]string                                              `json:"affected_channels,omitempty"`
	AffectedEntity           *enums.ProductcatalogdiagnosticgroupAffectedEntity     `json:"affected_entity,omitempty"`
	AffectedFeatures         *[]enums.ProductcatalogdiagnosticgroupAffectedFeatures `json:"affected_features,omitempty"`
	Diagnostics              *[]map[string]interface{}                              `json:"diagnostics,omitempty"`
	ErrorCode                *int                                                   `json:"error_code,omitempty"`
	NumberOfAffectedEntities *int                                                   `json:"number_of_affected_entities,omitempty"`
	NumberOfAffectedItems    *int                                                   `json:"number_of_affected_items,omitempty"`
	Severity                 *enums.ProductcatalogdiagnosticgroupSeverity           `json:"severity,omitempty"`
	Subtitle                 *string                                                `json:"subtitle,omitempty"`
	Title                    *string                                                `json:"title,omitempty"`
	Type                     *enums.ProductcatalogdiagnosticgroupType               `json:"type,omitempty"`
}

var ProductCatalogDiagnosticGroupFields = struct {
	AffectedChannels         string
	AffectedEntity           string
	AffectedFeatures         string
	Diagnostics              string
	ErrorCode                string
	NumberOfAffectedEntities string
	NumberOfAffectedItems    string
	Severity                 string
	Subtitle                 string
	Title                    string
	Type                     string
}{
	AffectedChannels:         "affected_channels",
	AffectedEntity:           "affected_entity",
	AffectedFeatures:         "affected_features",
	Diagnostics:              "diagnostics",
	ErrorCode:                "error_code",
	NumberOfAffectedEntities: "number_of_affected_entities",
	NumberOfAffectedItems:    "number_of_affected_items",
	Severity:                 "severity",
	Subtitle:                 "subtitle",
	Title:                    "title",
	Type:                     "type",
}
