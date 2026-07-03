package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetCanvasPreviewParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCanvasPreviewParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCanvasPreview(ctx context.Context, client *core.Client, id string, params GetCanvasPreviewParams) (*core.Cursor[objects.CanvasPreview], error) {
	var out core.Cursor[objects.CanvasPreview]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "preview"), params.ToParams(), &out)
	return &out, err
}

type GetCanvasPreviewsParams struct {
	UserIds *[]core.ID  `facebook:"user_ids"`
	Extra   core.Params `facebook:"-"`
}

func (params GetCanvasPreviewsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.UserIds != nil {
		out["user_ids"] = *params.UserIds
	}
	return out
}

func GetCanvasPreviews(ctx context.Context, client *core.Client, id string, params GetCanvasPreviewsParams) (*core.Cursor[objects.TextWithEntities], error) {
	var out core.Cursor[objects.TextWithEntities]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "previews"), params.ToParams(), &out)
	return &out, err
}

type GetCanvasParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCanvasParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCanvas(ctx context.Context, client *core.Client, id string, params GetCanvasParams) (*objects.Canvas, error) {
	var out objects.Canvas
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateCanvasParams struct {
	BackgroundColor           *string     `facebook:"background_color"`
	BodyElementIds            *[]core.ID  `facebook:"body_element_ids"`
	EnableSwipeToOpen         *bool       `facebook:"enable_swipe_to_open"`
	HeroAssetFacebookPostID   *core.ID    `facebook:"hero_asset_facebook_post_id"`
	HeroAssetInstagramMediaID *core.ID    `facebook:"hero_asset_instagram_media_id"`
	IsHidden                  *bool       `facebook:"is_hidden"`
	IsPublished               *bool       `facebook:"is_published"`
	Name                      *string     `facebook:"name"`
	SourceTemplateID          *core.ID    `facebook:"source_template_id"`
	Extra                     core.Params `facebook:"-"`
}

func (params UpdateCanvasParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BackgroundColor != nil {
		out["background_color"] = *params.BackgroundColor
	}
	if params.BodyElementIds != nil {
		out["body_element_ids"] = *params.BodyElementIds
	}
	if params.EnableSwipeToOpen != nil {
		out["enable_swipe_to_open"] = *params.EnableSwipeToOpen
	}
	if params.HeroAssetFacebookPostID != nil {
		out["hero_asset_facebook_post_id"] = *params.HeroAssetFacebookPostID
	}
	if params.HeroAssetInstagramMediaID != nil {
		out["hero_asset_instagram_media_id"] = *params.HeroAssetInstagramMediaID
	}
	if params.IsHidden != nil {
		out["is_hidden"] = *params.IsHidden
	}
	if params.IsPublished != nil {
		out["is_published"] = *params.IsPublished
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.SourceTemplateID != nil {
		out["source_template_id"] = *params.SourceTemplateID
	}
	return out
}

func UpdateCanvas(ctx context.Context, client *core.Client, id string, params UpdateCanvasParams) (*objects.Canvas, error) {
	var out objects.Canvas
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
