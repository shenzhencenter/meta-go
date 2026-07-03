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

func GetImageCopyright(ctx context.Context, client *core.Client, id string, params GetImageCopyrightParams) (*objects.ImageCopyright, error) {
	var out objects.ImageCopyright
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func UpdateImageCopyright(ctx context.Context, client *core.Client, id string, params UpdateImageCopyrightParams) (*objects.ImageCopyright, error) {
	var out objects.ImageCopyright
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
