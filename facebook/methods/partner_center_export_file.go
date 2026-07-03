package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPartnerCenterExportFileParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPartnerCenterExportFileParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPartnerCenterExportFile(ctx context.Context, client *core.Client, id string, params GetPartnerCenterExportFileParams) (*objects.PartnerCenterExportFile, error) {
	var out objects.PartnerCenterExportFile
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
