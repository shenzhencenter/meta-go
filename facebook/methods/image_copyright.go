package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetImageCopyrightParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetImageCopyrightParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetImageCopyrightBatchCall(id string, params GetImageCopyrightParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetImageCopyrightBatchRequest(id string, params GetImageCopyrightParams, options ...core.BatchOption) *core.BatchRequest[objects.ImageCopyright] {
	return core.NewBatchRequest[objects.ImageCopyright](GetImageCopyrightBatchCall(id, params, options...))
}

func DecodeGetImageCopyrightBatchResponse(response *core.BatchResponse) (*objects.ImageCopyright, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ImageCopyright
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetImageCopyrightWithResponse(ctx context.Context, client *core.Client, id string, params GetImageCopyrightParams) (*objects.ImageCopyright, *core.Response, error) {
	var out objects.ImageCopyright
	call := GetImageCopyrightBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetImageCopyright(ctx context.Context, client *core.Client, id string, params GetImageCopyrightParams) (*objects.ImageCopyright, error) {
	out, _, err := GetImageCopyrightWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateImageCopyrightParams struct {
	Artist                      *string                             `facebook:"artist"`
	Creator                     *string                             `facebook:"creator"`
	CustomID                    *core.ID                            `facebook:"custom_id"`
	Description                 *string                             `facebook:"description"`
	GeoOwnership                *[]enums.ImagecopyrightGeoOwnership `facebook:"geo_ownership"`
	OriginalContentCreationDate *uint64                             `facebook:"original_content_creation_date"`
	Title                       *string                             `facebook:"title"`
	Extra                       core.Params                         `facebook:"-"`
}

func (params UpdateImageCopyrightParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Artist != nil {
		out["artist"] = *params.Artist
	}
	if params.Creator != nil {
		out["creator"] = *params.Creator
	}
	if params.CustomID != nil {
		out["custom_id"] = *params.CustomID
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.GeoOwnership != nil {
		out["geo_ownership"] = *params.GeoOwnership
	}
	if params.OriginalContentCreationDate != nil {
		out["original_content_creation_date"] = *params.OriginalContentCreationDate
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	return out
}

func UpdateImageCopyrightBatchCall(id string, params UpdateImageCopyrightParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateImageCopyrightBatchRequest(id string, params UpdateImageCopyrightParams, options ...core.BatchOption) *core.BatchRequest[objects.ImageCopyright] {
	return core.NewBatchRequest[objects.ImageCopyright](UpdateImageCopyrightBatchCall(id, params, options...))
}

func DecodeUpdateImageCopyrightBatchResponse(response *core.BatchResponse) (*objects.ImageCopyright, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ImageCopyright
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateImageCopyrightWithResponse(ctx context.Context, client *core.Client, id string, params UpdateImageCopyrightParams) (*objects.ImageCopyright, *core.Response, error) {
	var out objects.ImageCopyright
	call := UpdateImageCopyrightBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateImageCopyright(ctx context.Context, client *core.Client, id string, params UpdateImageCopyrightParams) (*objects.ImageCopyright, error) {
	out, _, err := UpdateImageCopyrightWithResponse(ctx, client, id, params)
	return out, err
}
