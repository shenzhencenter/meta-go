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
