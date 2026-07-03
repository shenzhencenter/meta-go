package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CatalogWebsiteOnboardingSettings struct {
	ID          *core.ID `json:"id,omitempty"`
	QualityBand *string  `json:"quality_band,omitempty"`
	Status      *string  `json:"status,omitempty"`
}
