package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteOpenBridgeConfigurationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteOpenBridgeConfigurationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteOpenBridgeConfigurationBatchCall(id string, params DeleteOpenBridgeConfigurationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteOpenBridgeConfigurationBatchRequest(id string, params DeleteOpenBridgeConfigurationParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteOpenBridgeConfigurationBatchCall(id, params, options...))
}

func DecodeDeleteOpenBridgeConfigurationBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteOpenBridgeConfiguration(ctx context.Context, client *core.Client, id string, params DeleteOpenBridgeConfigurationParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteOpenBridgeConfigurationBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetOpenBridgeConfigurationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOpenBridgeConfigurationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOpenBridgeConfigurationBatchCall(id string, params GetOpenBridgeConfigurationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetOpenBridgeConfigurationBatchRequest(id string, params GetOpenBridgeConfigurationParams, options ...core.BatchOption) *core.BatchRequest[objects.OpenBridgeConfiguration] {
	return core.NewBatchRequest[objects.OpenBridgeConfiguration](GetOpenBridgeConfigurationBatchCall(id, params, options...))
}

func DecodeGetOpenBridgeConfigurationBatchResponse(response *core.BatchResponse) (*objects.OpenBridgeConfiguration, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.OpenBridgeConfiguration
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOpenBridgeConfiguration(ctx context.Context, client *core.Client, id string, params GetOpenBridgeConfigurationParams) (*objects.OpenBridgeConfiguration, error) {
	var out objects.OpenBridgeConfiguration
	call := GetOpenBridgeConfigurationBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateOpenBridgeConfigurationParams struct {
	Active                         *bool                                                        `facebook:"active"`
	BlockedEventTypes              *[]string                                                    `facebook:"blocked_event_types"`
	BlockedWebsites                *[]string                                                    `facebook:"blocked_websites"`
	CapiPublishingState            *enums.OpenbridgeconfigurationCapiPublishingState            `facebook:"capi_publishing_state"`
	CloudProvider                  *string                                                      `facebook:"cloud_provider"`
	CloudRegion                    *string                                                      `facebook:"cloud_region"`
	DestinationID                  *core.ID                                                     `facebook:"destination_id"`
	Endpoint                       *string                                                      `facebook:"endpoint"`
	EventEnrichmentAdvertiserState *enums.OpenbridgeconfigurationEventEnrichmentAdvertiserState `facebook:"event_enrichment_advertiser_state"`
	EventEnrichmentMetaState       *enums.OpenbridgeconfigurationEventEnrichmentMetaState       `facebook:"event_enrichment_meta_state"`
	EventEnrichmentState           *enums.OpenbridgeconfigurationEventEnrichmentState           `facebook:"event_enrichment_state"`
	FallbackDomain                 *string                                                      `facebook:"fallback_domain"`
	HostBusinessID                 *core.ID                                                     `facebook:"host_business_id"`
	InstanceID                     *core.ID                                                     `facebook:"instance_id"`
	InstanceVersion                *string                                                      `facebook:"instance_version"`
	IsSgwInstance                  *bool                                                        `facebook:"is_sgw_instance"`
	IsSgwPixelFromMetaPixel        *bool                                                        `facebook:"is_sgw_pixel_from_meta_pixel"`
	PartnerName                    *string                                                      `facebook:"partner_name"`
	SgwAccountID                   *core.ID                                                     `facebook:"sgw_account_id"`
	SgwInstanceURL                 *string                                                      `facebook:"sgw_instance_url"`
	SgwPixelID                     *core.ID                                                     `facebook:"sgw_pixel_id"`
	Extra                          core.Params                                                  `facebook:"-"`
}

func (params UpdateOpenBridgeConfigurationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Active != nil {
		out["active"] = *params.Active
	}
	if params.BlockedEventTypes != nil {
		out["blocked_event_types"] = *params.BlockedEventTypes
	}
	if params.BlockedWebsites != nil {
		out["blocked_websites"] = *params.BlockedWebsites
	}
	if params.CapiPublishingState != nil {
		out["capi_publishing_state"] = *params.CapiPublishingState
	}
	if params.CloudProvider != nil {
		out["cloud_provider"] = *params.CloudProvider
	}
	if params.CloudRegion != nil {
		out["cloud_region"] = *params.CloudRegion
	}
	if params.DestinationID != nil {
		out["destination_id"] = *params.DestinationID
	}
	if params.Endpoint != nil {
		out["endpoint"] = *params.Endpoint
	}
	if params.EventEnrichmentAdvertiserState != nil {
		out["event_enrichment_advertiser_state"] = *params.EventEnrichmentAdvertiserState
	}
	if params.EventEnrichmentMetaState != nil {
		out["event_enrichment_meta_state"] = *params.EventEnrichmentMetaState
	}
	if params.EventEnrichmentState != nil {
		out["event_enrichment_state"] = *params.EventEnrichmentState
	}
	if params.FallbackDomain != nil {
		out["fallback_domain"] = *params.FallbackDomain
	}
	if params.HostBusinessID != nil {
		out["host_business_id"] = *params.HostBusinessID
	}
	if params.InstanceID != nil {
		out["instance_id"] = *params.InstanceID
	}
	if params.InstanceVersion != nil {
		out["instance_version"] = *params.InstanceVersion
	}
	if params.IsSgwInstance != nil {
		out["is_sgw_instance"] = *params.IsSgwInstance
	}
	if params.IsSgwPixelFromMetaPixel != nil {
		out["is_sgw_pixel_from_meta_pixel"] = *params.IsSgwPixelFromMetaPixel
	}
	if params.PartnerName != nil {
		out["partner_name"] = *params.PartnerName
	}
	if params.SgwAccountID != nil {
		out["sgw_account_id"] = *params.SgwAccountID
	}
	if params.SgwInstanceURL != nil {
		out["sgw_instance_url"] = *params.SgwInstanceURL
	}
	if params.SgwPixelID != nil {
		out["sgw_pixel_id"] = *params.SgwPixelID
	}
	return out
}

func UpdateOpenBridgeConfigurationBatchCall(id string, params UpdateOpenBridgeConfigurationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateOpenBridgeConfigurationBatchRequest(id string, params UpdateOpenBridgeConfigurationParams, options ...core.BatchOption) *core.BatchRequest[objects.OpenBridgeConfiguration] {
	return core.NewBatchRequest[objects.OpenBridgeConfiguration](UpdateOpenBridgeConfigurationBatchCall(id, params, options...))
}

func DecodeUpdateOpenBridgeConfigurationBatchResponse(response *core.BatchResponse) (*objects.OpenBridgeConfiguration, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.OpenBridgeConfiguration
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateOpenBridgeConfiguration(ctx context.Context, client *core.Client, id string, params UpdateOpenBridgeConfigurationParams) (*objects.OpenBridgeConfiguration, error) {
	var out objects.OpenBridgeConfiguration
	call := UpdateOpenBridgeConfigurationBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
