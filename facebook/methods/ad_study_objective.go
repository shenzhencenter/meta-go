package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdStudyObjectiveAdspixelsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyObjectiveAdspixelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyObjectiveAdspixelsBatchCall(id string, params GetAdStudyObjectiveAdspixelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adspixels"), params.ToParams(), options...)
}

func NewGetAdStudyObjectiveAdspixelsBatchRequest(id string, params GetAdStudyObjectiveAdspixelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsPixel]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsPixel]](GetAdStudyObjectiveAdspixelsBatchCall(id, params, options...))
}

func DecodeGetAdStudyObjectiveAdspixelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsPixel], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsPixel]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyObjectiveAdspixels(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectiveAdspixelsParams) (*core.Cursor[objects.AdsPixel], error) {
	var out core.Cursor[objects.AdsPixel]
	call := GetAdStudyObjectiveAdspixelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdStudyObjectiveApplicationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyObjectiveApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyObjectiveApplicationsBatchCall(id string, params GetAdStudyObjectiveApplicationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "applications"), params.ToParams(), options...)
}

func NewGetAdStudyObjectiveApplicationsBatchRequest(id string, params GetAdStudyObjectiveApplicationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Application]] {
	return core.NewBatchRequest[core.Cursor[objects.Application]](GetAdStudyObjectiveApplicationsBatchCall(id, params, options...))
}

func DecodeGetAdStudyObjectiveApplicationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Application], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Application]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyObjectiveApplications(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectiveApplicationsParams) (*core.Cursor[objects.Application], error) {
	var out core.Cursor[objects.Application]
	call := GetAdStudyObjectiveApplicationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdStudyObjectiveBrandRequestsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyObjectiveBrandRequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyObjectiveBrandRequestsBatchCall(id string, params GetAdStudyObjectiveBrandRequestsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "brand_requests"), params.ToParams(), options...)
}

func NewGetAdStudyObjectiveBrandRequestsBatchRequest(id string, params GetAdStudyObjectiveBrandRequestsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BrandRequest]] {
	return core.NewBatchRequest[core.Cursor[objects.BrandRequest]](GetAdStudyObjectiveBrandRequestsBatchCall(id, params, options...))
}

func DecodeGetAdStudyObjectiveBrandRequestsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BrandRequest], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BrandRequest]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyObjectiveBrandRequests(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectiveBrandRequestsParams) (*core.Cursor[objects.BrandRequest], error) {
	var out core.Cursor[objects.BrandRequest]
	call := GetAdStudyObjectiveBrandRequestsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdStudyObjectiveCustomconversionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyObjectiveCustomconversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyObjectiveCustomconversionsBatchCall(id string, params GetAdStudyObjectiveCustomconversionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "customconversions"), params.ToParams(), options...)
}

func NewGetAdStudyObjectiveCustomconversionsBatchRequest(id string, params GetAdStudyObjectiveCustomconversionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CustomConversion]] {
	return core.NewBatchRequest[core.Cursor[objects.CustomConversion]](GetAdStudyObjectiveCustomconversionsBatchCall(id, params, options...))
}

func DecodeGetAdStudyObjectiveCustomconversionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CustomConversion], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CustomConversion]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyObjectiveCustomconversions(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectiveCustomconversionsParams) (*core.Cursor[objects.CustomConversion], error) {
	var out core.Cursor[objects.CustomConversion]
	call := GetAdStudyObjectiveCustomconversionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdStudyObjectiveOfflineConversionDataSetsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyObjectiveOfflineConversionDataSetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyObjectiveOfflineConversionDataSetsBatchCall(id string, params GetAdStudyObjectiveOfflineConversionDataSetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "offline_conversion_data_sets"), params.ToParams(), options...)
}

func NewGetAdStudyObjectiveOfflineConversionDataSetsBatchRequest(id string, params GetAdStudyObjectiveOfflineConversionDataSetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OfflineConversionDataSet]] {
	return core.NewBatchRequest[core.Cursor[objects.OfflineConversionDataSet]](GetAdStudyObjectiveOfflineConversionDataSetsBatchCall(id, params, options...))
}

func DecodeGetAdStudyObjectiveOfflineConversionDataSetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OfflineConversionDataSet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.OfflineConversionDataSet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyObjectiveOfflineConversionDataSets(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectiveOfflineConversionDataSetsParams) (*core.Cursor[objects.OfflineConversionDataSet], error) {
	var out core.Cursor[objects.OfflineConversionDataSet]
	call := GetAdStudyObjectiveOfflineConversionDataSetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdStudyObjectivePartnerPrivateStudiesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyObjectivePartnerPrivateStudiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyObjectivePartnerPrivateStudiesBatchCall(id string, params GetAdStudyObjectivePartnerPrivateStudiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "partner_private_studies"), params.ToParams(), options...)
}

func NewGetAdStudyObjectivePartnerPrivateStudiesBatchRequest(id string, params GetAdStudyObjectivePartnerPrivateStudiesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetAdStudyObjectivePartnerPrivateStudiesBatchCall(id, params, options...))
}

func DecodeGetAdStudyObjectivePartnerPrivateStudiesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Business]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyObjectivePartnerPrivateStudies(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectivePartnerPrivateStudiesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	call := GetAdStudyObjectivePartnerPrivateStudiesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdStudyObjectivePartnerstudiesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyObjectivePartnerstudiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyObjectivePartnerstudiesBatchCall(id string, params GetAdStudyObjectivePartnerstudiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "partnerstudies"), params.ToParams(), options...)
}

func NewGetAdStudyObjectivePartnerstudiesBatchRequest(id string, params GetAdStudyObjectivePartnerstudiesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PartnerStudy]] {
	return core.NewBatchRequest[core.Cursor[objects.PartnerStudy]](GetAdStudyObjectivePartnerstudiesBatchCall(id, params, options...))
}

func DecodeGetAdStudyObjectivePartnerstudiesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PartnerStudy], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PartnerStudy]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyObjectivePartnerstudies(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectivePartnerstudiesParams) (*core.Cursor[objects.PartnerStudy], error) {
	var out core.Cursor[objects.PartnerStudy]
	call := GetAdStudyObjectivePartnerstudiesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdStudyObjectiveParams struct {
	Breakdowns *[]enums.AdstudyobjectiveBreakdowns `facebook:"breakdowns"`
	Ds         *string                             `facebook:"ds"`
	Extra      core.Params                         `facebook:"-"`
}

func (params GetAdStudyObjectiveParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Breakdowns != nil {
		out["breakdowns"] = *params.Breakdowns
	}
	if params.Ds != nil {
		out["ds"] = *params.Ds
	}
	return out
}

func GetAdStudyObjectiveBatchCall(id string, params GetAdStudyObjectiveParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdStudyObjectiveBatchRequest(id string, params GetAdStudyObjectiveParams, options ...core.BatchOption) *core.BatchRequest[objects.AdStudyObjective] {
	return core.NewBatchRequest[objects.AdStudyObjective](GetAdStudyObjectiveBatchCall(id, params, options...))
}

func DecodeGetAdStudyObjectiveBatchResponse(response *core.BatchResponse) (*objects.AdStudyObjective, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdStudyObjective
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyObjective(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectiveParams) (*objects.AdStudyObjective, error) {
	var out objects.AdStudyObjective
	call := GetAdStudyObjectiveBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateAdStudyObjectiveParams struct {
	Adspixels                 *[]map[string]interface{}   `facebook:"adspixels"`
	Applications              *[]map[string]interface{}   `facebook:"applications"`
	Customconversions         *[]map[string]interface{}   `facebook:"customconversions"`
	IsPrimary                 *bool                       `facebook:"is_primary"`
	Name                      *string                     `facebook:"name"`
	OfflineConversionDataSets *[]map[string]interface{}   `facebook:"offline_conversion_data_sets"`
	OffsiteDatasets           *[]map[string]interface{}   `facebook:"offsite_datasets"`
	ProductCatalogs           *[]map[string]interface{}   `facebook:"product_catalogs"`
	ProductSets               *[]map[string]interface{}   `facebook:"product_sets"`
	Type                      *enums.AdstudyobjectiveType `facebook:"type"`
	Extra                     core.Params                 `facebook:"-"`
}

func (params UpdateAdStudyObjectiveParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Adspixels != nil {
		out["adspixels"] = *params.Adspixels
	}
	if params.Applications != nil {
		out["applications"] = *params.Applications
	}
	if params.Customconversions != nil {
		out["customconversions"] = *params.Customconversions
	}
	if params.IsPrimary != nil {
		out["is_primary"] = *params.IsPrimary
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.OfflineConversionDataSets != nil {
		out["offline_conversion_data_sets"] = *params.OfflineConversionDataSets
	}
	if params.OffsiteDatasets != nil {
		out["offsite_datasets"] = *params.OffsiteDatasets
	}
	if params.ProductCatalogs != nil {
		out["product_catalogs"] = *params.ProductCatalogs
	}
	if params.ProductSets != nil {
		out["product_sets"] = *params.ProductSets
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func UpdateAdStudyObjectiveBatchCall(id string, params UpdateAdStudyObjectiveParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateAdStudyObjectiveBatchRequest(id string, params UpdateAdStudyObjectiveParams, options ...core.BatchOption) *core.BatchRequest[objects.AdStudyObjective] {
	return core.NewBatchRequest[objects.AdStudyObjective](UpdateAdStudyObjectiveBatchCall(id, params, options...))
}

func DecodeUpdateAdStudyObjectiveBatchResponse(response *core.BatchResponse) (*objects.AdStudyObjective, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdStudyObjective
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateAdStudyObjective(ctx context.Context, client *core.Client, id string, params UpdateAdStudyObjectiveParams) (*objects.AdStudyObjective, error) {
	var out objects.AdStudyObjective
	call := UpdateAdStudyObjectiveBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
