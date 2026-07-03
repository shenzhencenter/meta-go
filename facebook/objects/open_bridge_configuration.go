package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type OpenBridgeConfiguration struct {
	Active                         *bool     `json:"active,omitempty"`
	BlockedEventTypes              *[]string `json:"blocked_event_types,omitempty"`
	BlockedWebsites                *[]string `json:"blocked_websites,omitempty"`
	BrowserAgent                   *[]string `json:"browser_agent,omitempty"`
	CapiPublishingState            *string   `json:"capi_publishing_state,omitempty"`
	CloudProvider                  *string   `json:"cloud_provider,omitempty"`
	CloudRegion                    *string   `json:"cloud_region,omitempty"`
	DestinationID                  *core.ID  `json:"destination_id,omitempty"`
	Endpoint                       *string   `json:"endpoint,omitempty"`
	EventEnrichmentAdvertiserState *string   `json:"event_enrichment_advertiser_state,omitempty"`
	EventEnrichmentMetaState       *string   `json:"event_enrichment_meta_state,omitempty"`
	EventEnrichmentState           *string   `json:"event_enrichment_state,omitempty"`
	FallbackDomain                 *string   `json:"fallback_domain,omitempty"`
	HostBusinessID                 *core.ID  `json:"host_business_id,omitempty"`
	ID                             *core.ID  `json:"id,omitempty"`
	InstanceID                     *core.ID  `json:"instance_id,omitempty"`
	InstanceVersion                *string   `json:"instance_version,omitempty"`
	IsSgwInstance                  *bool     `json:"is_sgw_instance,omitempty"`
	IsSgwPixelFromMetaPixel        *bool     `json:"is_sgw_pixel_from_meta_pixel,omitempty"`
	PartnerName                    *string   `json:"partner_name,omitempty"`
	PixelID                        *core.ID  `json:"pixel_id,omitempty"`
	SgwAccountID                   *core.ID  `json:"sgw_account_id,omitempty"`
	SgwInstanceURL                 *string   `json:"sgw_instance_url,omitempty"`
	SgwPixelID                     *core.ID  `json:"sgw_pixel_id,omitempty"`
}
