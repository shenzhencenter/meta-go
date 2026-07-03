package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetAdStudyObjectiveAdspixels(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectiveAdspixelsParams) (*core.Cursor[objects.AdsPixel], error) {
	var out core.Cursor[objects.AdsPixel]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adspixels"), params.ToParams(), &out)
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

func GetAdStudyObjectiveApplications(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectiveApplicationsParams) (*core.Cursor[objects.Application], error) {
	var out core.Cursor[objects.Application]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "applications"), params.ToParams(), &out)
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

func GetAdStudyObjectiveBrandRequests(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectiveBrandRequestsParams) (*core.Cursor[objects.BrandRequest], error) {
	var out core.Cursor[objects.BrandRequest]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "brand_requests"), params.ToParams(), &out)
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

func GetAdStudyObjectiveCustomconversions(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectiveCustomconversionsParams) (*core.Cursor[objects.CustomConversion], error) {
	var out core.Cursor[objects.CustomConversion]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "customconversions"), params.ToParams(), &out)
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

func GetAdStudyObjectiveOfflineConversionDataSets(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectiveOfflineConversionDataSetsParams) (*core.Cursor[objects.OfflineConversionDataSet], error) {
	var out core.Cursor[objects.OfflineConversionDataSet]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "offline_conversion_data_sets"), params.ToParams(), &out)
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

func GetAdStudyObjectivePartnerPrivateStudies(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectivePartnerPrivateStudiesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "partner_private_studies"), params.ToParams(), &out)
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

func GetAdStudyObjectivePartnerstudies(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectivePartnerstudiesParams) (*core.Cursor[objects.PartnerStudy], error) {
	var out core.Cursor[objects.PartnerStudy]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "partnerstudies"), params.ToParams(), &out)
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

func GetAdStudyObjective(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectiveParams) (*objects.AdStudyObjective, error) {
	var out objects.AdStudyObjective
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
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

func UpdateAdStudyObjective(ctx context.Context, client *core.Client, id string, params UpdateAdStudyObjectiveParams) (*objects.AdStudyObjective, error) {
	var out objects.AdStudyObjective
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
