package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func CreateAdCreativeAdlabels(ctx context.Context, client *core.Client, id string, params CreateAdCreativeAdlabelsParams) (*objects.AdCreative, error) {
	var out objects.AdCreative
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "adlabels"), params.ToParams(), &out)
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

func GetAdCreativeCreativeInsights(ctx context.Context, client *core.Client, id string, params GetAdCreativeCreativeInsightsParams) (*core.Cursor[objects.AdCreativeInsights], error) {
	var out core.Cursor[objects.AdCreativeInsights]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "creative_insights"), params.ToParams(), &out)
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

func GetAdCreativePreviews(ctx context.Context, client *core.Client, id string, params GetAdCreativePreviewsParams) (*core.Cursor[objects.AdPreview], error) {
	var out core.Cursor[objects.AdPreview]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "previews"), params.ToParams(), &out)
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

func DeleteAdCreative(ctx context.Context, client *core.Client, id string, params DeleteAdCreativeParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
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

func GetAdCreative(ctx context.Context, client *core.Client, id string, params GetAdCreativeParams) (*objects.AdCreative, error) {
	var out objects.AdCreative
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
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

func UpdateAdCreative(ctx context.Context, client *core.Client, id string, params UpdateAdCreativeParams) (*objects.AdCreative, error) {
	var out objects.AdCreative
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
