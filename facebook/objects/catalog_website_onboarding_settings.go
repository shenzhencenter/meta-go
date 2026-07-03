package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CatalogWebsiteOnboardingSettings struct {
	ID          *core.ID `json:"id,omitempty"`
	QualityBand *string  `json:"quality_band,omitempty"`
	Status      *string  `json:"status,omitempty"`
}
