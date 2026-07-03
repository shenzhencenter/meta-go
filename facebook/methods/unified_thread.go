package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetUnifiedThreadMessagesParams struct {
	Source *enums.UnifiedthreadmessagesSourceEnumParam `facebook:"source"`
	Extra  core.Params                                 `facebook:"-"`
}

func (params GetUnifiedThreadMessagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	return out
}

func GetUnifiedThreadMessages(ctx context.Context, client *core.Client, id string, params GetUnifiedThreadMessagesParams) (*core.Cursor[objects.UnifiedMessage], error) {
	var out core.Cursor[objects.UnifiedMessage]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "messages"), params.ToParams(), &out)
	return &out, err
}

type GetUnifiedThreadParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUnifiedThreadParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUnifiedThread(ctx context.Context, client *core.Client, id string, params GetUnifiedThreadParams) (*objects.UnifiedThread, error) {
	var out objects.UnifiedThread
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
