package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdgroupFacebookFeedbackCommentsParams struct {
	Order *enums.AdgroupfacebookfeedbackcommentsOrderEnumParam `facebook:"order"`
	Extra core.Params                                          `facebook:"-"`
}

func (params GetAdgroupFacebookFeedbackCommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Order != nil {
		out["order"] = *params.Order
	}
	return out
}

func GetAdgroupFacebookFeedbackComments(ctx context.Context, client *core.Client, id string, params GetAdgroupFacebookFeedbackCommentsParams) (*core.Cursor[objects.Comment], error) {
	var out core.Cursor[objects.Comment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
}
