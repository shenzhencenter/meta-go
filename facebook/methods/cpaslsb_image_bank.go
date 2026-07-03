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

func GetCPASLsbImageBankBackupImages(ctx context.Context, client *core.Client, id string, params GetCPASLsbImageBankBackupImagesParams) (*core.Cursor[objects.ProductImage], error) {
	var out core.Cursor[objects.ProductImage]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "backup_images"), params.ToParams(), &out)
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

func GetCPASLsbImageBank(ctx context.Context, client *core.Client, id string, params GetCPASLsbImageBankParams) (*objects.CPASLsbImageBank, error) {
	var out objects.CPASLsbImageBank
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
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

func UpdateCPASLsbImageBank(ctx context.Context, client *core.Client, id string, params UpdateCPASLsbImageBankParams) (*objects.CPASLsbImageBank, error) {
	var out objects.CPASLsbImageBank
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
