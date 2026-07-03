package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteAdgroupGendeleteParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdgroupGendeleteParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdgroupGendelete(ctx context.Context, client *core.Client, id string, params DeleteAdgroupGendeleteParams) (*objects.AdgroupDelete, error) {
	var out objects.AdgroupDelete
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, ""), params.ToParams(), &out)
	return &out, err
}
