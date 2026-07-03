package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type CreatePublisherBlockListAppendPublisherUrlsParams struct {
	PublisherUrls []string    `facebook:"publisher_urls"`
	Extra         core.Params `facebook:"-"`
}

func (params CreatePublisherBlockListAppendPublisherUrlsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["publisher_urls"] = params.PublisherUrls
	return out
}

func CreatePublisherBlockListAppendPublisherUrls(ctx context.Context, client *core.Client, id string, params CreatePublisherBlockListAppendPublisherUrlsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "append_publisher_urls"), params.ToParams(), &out)
	return &out, err
}

type GetPublisherBlockListPagedWebPublishersParams struct {
	DraftID *core.ID    `facebook:"draft_id"`
	Extra   core.Params `facebook:"-"`
}

func (params GetPublisherBlockListPagedWebPublishersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DraftID != nil {
		out["draft_id"] = *params.DraftID
	}
	return out
}

func GetPublisherBlockListPagedWebPublishers(ctx context.Context, client *core.Client, id string, params GetPublisherBlockListPagedWebPublishersParams) (*core.Cursor[objects.WebPublisher], error) {
	var out core.Cursor[objects.WebPublisher]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "paged_web_publishers"), params.ToParams(), &out)
	return &out, err
}

type DeletePublisherBlockListParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeletePublisherBlockListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeletePublisherBlockList(ctx context.Context, client *core.Client, id string, params DeletePublisherBlockListParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetPublisherBlockListParams struct {
	AccountID  *core.ID    `facebook:"account_id"`
	BusinessID *core.ID    `facebook:"business_id"`
	DraftID    *core.ID    `facebook:"draft_id"`
	Extra      core.Params `facebook:"-"`
}

func (params GetPublisherBlockListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AccountID != nil {
		out["account_id"] = *params.AccountID
	}
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	if params.DraftID != nil {
		out["draft_id"] = *params.DraftID
	}
	return out
}

func GetPublisherBlockList(ctx context.Context, client *core.Client, id string, params GetPublisherBlockListParams) (*objects.PublisherBlockList, error) {
	var out objects.PublisherBlockList
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdatePublisherBlockListParams struct {
	Spec  map[string]interface{} `facebook:"spec"`
	Extra core.Params            `facebook:"-"`
}

func (params UpdatePublisherBlockListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["spec"] = params.Spec
	return out
}

func UpdatePublisherBlockList(ctx context.Context, client *core.Client, id string, params UpdatePublisherBlockListParams) (*objects.PublisherBlockList, error) {
	var out objects.PublisherBlockList
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
