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

var OpenBridgeConfigurationFields = struct {
	Active                         string
	BlockedEventTypes              string
	BlockedWebsites                string
	BrowserAgent                   string
	CapiPublishingState            string
	CloudProvider                  string
	CloudRegion                    string
	DestinationID                  string
	Endpoint                       string
	EventEnrichmentAdvertiserState string
	EventEnrichmentMetaState       string
	EventEnrichmentState           string
	FallbackDomain                 string
	HostBusinessID                 string
	ID                             string
	InstanceID                     string
	InstanceVersion                string
	IsSgwInstance                  string
	IsSgwPixelFromMetaPixel        string
	PartnerName                    string
	PixelID                        string
	SgwAccountID                   string
	SgwInstanceURL                 string
	SgwPixelID                     string
}{
	Active:                         "active",
	BlockedEventTypes:              "blocked_event_types",
	BlockedWebsites:                "blocked_websites",
	BrowserAgent:                   "browser_agent",
	CapiPublishingState:            "capi_publishing_state",
	CloudProvider:                  "cloud_provider",
	CloudRegion:                    "cloud_region",
	DestinationID:                  "destination_id",
	Endpoint:                       "endpoint",
	EventEnrichmentAdvertiserState: "event_enrichment_advertiser_state",
	EventEnrichmentMetaState:       "event_enrichment_meta_state",
	EventEnrichmentState:           "event_enrichment_state",
	FallbackDomain:                 "fallback_domain",
	HostBusinessID:                 "host_business_id",
	ID:                             "id",
	InstanceID:                     "instance_id",
	InstanceVersion:                "instance_version",
	IsSgwInstance:                  "is_sgw_instance",
	IsSgwPixelFromMetaPixel:        "is_sgw_pixel_from_meta_pixel",
	PartnerName:                    "partner_name",
	PixelID:                        "pixel_id",
	SgwAccountID:                   "sgw_account_id",
	SgwInstanceURL:                 "sgw_instance_url",
	SgwPixelID:                     "sgw_pixel_id",
}
