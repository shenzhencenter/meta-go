package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCPASLsbImageBankBackupImagesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCPASLsbImageBankBackupImagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCPASLsbImageBankBackupImagesBatchCall(id string, params GetCPASLsbImageBankBackupImagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "backup_images"), params.ToParams(), options...)
}

func NewGetCPASLsbImageBankBackupImagesBatchRequest(id string, params GetCPASLsbImageBankBackupImagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductImage]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductImage]](GetCPASLsbImageBankBackupImagesBatchCall(id, params, options...))
}

func DecodeGetCPASLsbImageBankBackupImagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductImage], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductImage]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCPASLsbImageBankBackupImages(ctx context.Context, client *core.Client, id string, params GetCPASLsbImageBankBackupImagesParams) (*core.Cursor[objects.ProductImage], error) {
	var out core.Cursor[objects.ProductImage]
	call := GetCPASLsbImageBankBackupImagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetCPASLsbImageBankParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCPASLsbImageBankParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCPASLsbImageBankBatchCall(id string, params GetCPASLsbImageBankParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCPASLsbImageBankBatchRequest(id string, params GetCPASLsbImageBankParams, options ...core.BatchOption) *core.BatchRequest[objects.CPASLsbImageBank] {
	return core.NewBatchRequest[objects.CPASLsbImageBank](GetCPASLsbImageBankBatchCall(id, params, options...))
}

func DecodeGetCPASLsbImageBankBatchResponse(response *core.BatchResponse) (*objects.CPASLsbImageBank, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CPASLsbImageBank
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCPASLsbImageBank(ctx context.Context, client *core.Client, id string, params GetCPASLsbImageBankParams) (*objects.CPASLsbImageBank, error) {
	var out objects.CPASLsbImageBank
	call := GetCPASLsbImageBankBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateCPASLsbImageBankParams struct {
	BackupImageUrls []string    `facebook:"backup_image_urls"`
	Extra           core.Params `facebook:"-"`
}

func (params UpdateCPASLsbImageBankParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["backup_image_urls"] = params.BackupImageUrls
	return out
}

func UpdateCPASLsbImageBankBatchCall(id string, params UpdateCPASLsbImageBankParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateCPASLsbImageBankBatchRequest(id string, params UpdateCPASLsbImageBankParams, options ...core.BatchOption) *core.BatchRequest[objects.CPASLsbImageBank] {
	return core.NewBatchRequest[objects.CPASLsbImageBank](UpdateCPASLsbImageBankBatchCall(id, params, options...))
}

func DecodeUpdateCPASLsbImageBankBatchResponse(response *core.BatchResponse) (*objects.CPASLsbImageBank, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CPASLsbImageBank
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateCPASLsbImageBank(ctx context.Context, client *core.Client, id string, params UpdateCPASLsbImageBankParams) (*objects.CPASLsbImageBank, error) {
	var out objects.CPASLsbImageBank
	call := UpdateCPASLsbImageBankBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
