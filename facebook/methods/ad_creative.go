package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type CreateAdCreativeAdlabelsParams struct {
	Adlabels []map[string]interface{} `facebook:"adlabels"`
	Extra    core.Params              `facebook:"-"`
}

func (params CreateAdCreativeAdlabelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["adlabels"] = params.Adlabels
	return out
}

func CreateAdCreativeAdlabelsBatchCall(id string, params CreateAdCreativeAdlabelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adlabels"), params.ToParams(), options...)
}

func NewCreateAdCreativeAdlabelsBatchRequest(id string, params CreateAdCreativeAdlabelsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdCreative] {
	return core.NewBatchRequest[objects.AdCreative](CreateAdCreativeAdlabelsBatchCall(id, params, options...))
}

func DecodeCreateAdCreativeAdlabelsBatchResponse(response *core.BatchResponse) (*objects.AdCreative, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdCreative
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdCreativeAdlabels(ctx context.Context, client *core.Client, id string, params CreateAdCreativeAdlabelsParams) (*objects.AdCreative, error) {
	var out objects.AdCreative
	call := CreateAdCreativeAdlabelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdCreativeCreativeInsightsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdCreativeCreativeInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdCreativeCreativeInsightsBatchCall(id string, params GetAdCreativeCreativeInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "creative_insights"), params.ToParams(), options...)
}

func NewGetAdCreativeCreativeInsightsBatchRequest(id string, params GetAdCreativeCreativeInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdCreativeInsights]] {
	return core.NewBatchRequest[core.Cursor[objects.AdCreativeInsights]](GetAdCreativeCreativeInsightsBatchCall(id, params, options...))
}

func DecodeGetAdCreativeCreativeInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdCreativeInsights], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdCreativeInsights]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdCreativeCreativeInsights(ctx context.Context, client *core.Client, id string, params GetAdCreativeCreativeInsightsParams) (*core.Cursor[objects.AdCreativeInsights], error) {
	var out core.Cursor[objects.AdCreativeInsights]
	call := GetAdCreativeCreativeInsightsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdCreativePreviewsParams struct {
	AdFormat             enums.AdcreativepreviewsAdFormatEnumParam         `facebook:"ad_format"`
	CreativeFeature      *enums.AdcreativepreviewsCreativeFeatureEnumParam `facebook:"creative_feature"`
	DynamicAssetLabel    *string                                           `facebook:"dynamic_asset_label"`
	DynamicCreativeSpec  *map[string]interface{}                           `facebook:"dynamic_creative_spec"`
	DynamicCustomization *map[string]interface{}                           `facebook:"dynamic_customization"`
	EndDate              *time.Time                                        `facebook:"end_date"`
	Height               *uint64                                           `facebook:"height"`
	Locale               *string                                           `facebook:"locale"`
	PlacePageID          *core.ID                                          `facebook:"place_page_id"`
	Post                 *map[string]interface{}                           `facebook:"post"`
	ProductItemIds       *[]core.ID                                        `facebook:"product_item_ids"`
	RenderType           *enums.AdcreativepreviewsRenderTypeEnumParam      `facebook:"render_type"`
	StartDate            *time.Time                                        `facebook:"start_date"`
	Width                *uint64                                           `facebook:"width"`
	Extra                core.Params                                       `facebook:"-"`
}

func (params GetAdCreativePreviewsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["ad_format"] = params.AdFormat
	if params.CreativeFeature != nil {
		out["creative_feature"] = *params.CreativeFeature
	}
	if params.DynamicAssetLabel != nil {
		out["dynamic_asset_label"] = *params.DynamicAssetLabel
	}
	if params.DynamicCreativeSpec != nil {
		out["dynamic_creative_spec"] = *params.DynamicCreativeSpec
	}
	if params.DynamicCustomization != nil {
		out["dynamic_customization"] = *params.DynamicCustomization
	}
	if params.EndDate != nil {
		out["end_date"] = *params.EndDate
	}
	if params.Height != nil {
		out["height"] = *params.Height
	}
	if params.Locale != nil {
		out["locale"] = *params.Locale
	}
	if params.PlacePageID != nil {
		out["place_page_id"] = *params.PlacePageID
	}
	if params.Post != nil {
		out["post"] = *params.Post
	}
	if params.ProductItemIds != nil {
		out["product_item_ids"] = *params.ProductItemIds
	}
	if params.RenderType != nil {
		out["render_type"] = *params.RenderType
	}
	if params.StartDate != nil {
		out["start_date"] = *params.StartDate
	}
	if params.Width != nil {
		out["width"] = *params.Width
	}
	return out
}

func GetAdCreativePreviewsBatchCall(id string, params GetAdCreativePreviewsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "previews"), params.ToParams(), options...)
}

func NewGetAdCreativePreviewsBatchRequest(id string, params GetAdCreativePreviewsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdPreview]] {
	return core.NewBatchRequest[core.Cursor[objects.AdPreview]](GetAdCreativePreviewsBatchCall(id, params, options...))
}

func DecodeGetAdCreativePreviewsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdPreview], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdPreview]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdCreativePreviews(ctx context.Context, client *core.Client, id string, params GetAdCreativePreviewsParams) (*core.Cursor[objects.AdPreview], error) {
	var out core.Cursor[objects.AdPreview]
	call := GetAdCreativePreviewsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteAdCreativeParams struct {
	AccountID *core.ID                  `facebook:"account_id"`
	Adlabels  *[]map[string]interface{} `facebook:"adlabels"`
	Name      *string                   `facebook:"name"`
	Status    *enums.AdcreativeStatus   `facebook:"status"`
	Extra     core.Params               `facebook:"-"`
}

func (params DeleteAdCreativeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AccountID != nil {
		out["account_id"] = *params.AccountID
	}
	if params.Adlabels != nil {
		out["adlabels"] = *params.Adlabels
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	return out
}

func DeleteAdCreativeBatchCall(id string, params DeleteAdCreativeParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteAdCreativeBatchRequest(id string, params DeleteAdCreativeParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdCreativeBatchCall(id, params, options...))
}

func DecodeDeleteAdCreativeBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdCreative(ctx context.Context, client *core.Client, id string, params DeleteAdCreativeParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteAdCreativeBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdCreativeParams struct {
	ThumbnailHeight *uint64     `facebook:"thumbnail_height"`
	ThumbnailWidth  *uint64     `facebook:"thumbnail_width"`
	Extra           core.Params `facebook:"-"`
}

func (params GetAdCreativeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ThumbnailHeight != nil {
		out["thumbnail_height"] = *params.ThumbnailHeight
	}
	if params.ThumbnailWidth != nil {
		out["thumbnail_width"] = *params.ThumbnailWidth
	}
	return out
}

func GetAdCreativeBatchCall(id string, params GetAdCreativeParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdCreativeBatchRequest(id string, params GetAdCreativeParams, options ...core.BatchOption) *core.BatchRequest[objects.AdCreative] {
	return core.NewBatchRequest[objects.AdCreative](GetAdCreativeBatchCall(id, params, options...))
}

func DecodeGetAdCreativeBatchResponse(response *core.BatchResponse) (*objects.AdCreative, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdCreative
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdCreative(ctx context.Context, client *core.Client, id string, params GetAdCreativeParams) (*objects.AdCreative, error) {
	var out objects.AdCreative
	call := GetAdCreativeBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateAdCreativeParams struct {
	AccountID *core.ID                  `facebook:"account_id"`
	Adlabels  *[]map[string]interface{} `facebook:"adlabels"`
	Name      *string                   `facebook:"name"`
	Status    *enums.AdcreativeStatus   `facebook:"status"`
	Extra     core.Params               `facebook:"-"`
}

func (params UpdateAdCreativeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AccountID != nil {
		out["account_id"] = *params.AccountID
	}
	if params.Adlabels != nil {
		out["adlabels"] = *params.Adlabels
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	return out
}

func UpdateAdCreativeBatchCall(id string, params UpdateAdCreativeParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateAdCreativeBatchRequest(id string, params UpdateAdCreativeParams, options ...core.BatchOption) *core.BatchRequest[objects.AdCreative] {
	return core.NewBatchRequest[objects.AdCreative](UpdateAdCreativeBatchCall(id, params, options...))
}

func DecodeUpdateAdCreativeBatchResponse(response *core.BatchResponse) (*objects.AdCreative, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdCreative
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateAdCreative(ctx context.Context, client *core.Client, id string, params UpdateAdCreativeParams) (*objects.AdCreative, error) {
	var out objects.AdCreative
	call := UpdateAdCreativeBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
