package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type SignalsIWLExtractor struct {
	DomainURI       *string                 `json:"domain_uri,omitempty"`
	EventType       *string                 `json:"event_type,omitempty"`
	ExtractorConfig *map[string]interface{} `json:"extractor_config,omitempty"`
	ExtractorType   *string                 `json:"extractor_type,omitempty"`
	ID              *core.ID                `json:"id,omitempty"`
}
