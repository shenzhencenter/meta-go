package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetURLParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetURLParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetURL(ctx context.Context, client *core.Client, id string, params GetURLParams) (*objects.URL, error) {
	var out objects.URL
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateURLParams struct {
	Blacklist *bool              `facebook:"blacklist"`
	Denylist  *bool              `facebook:"denylist"`
	Hmac      *string            `facebook:"hmac"`
	Locale    *[]string          `facebook:"locale"`
	Scopes    *[]enums.URLScopes `facebook:"scopes"`
	Ts        *time.Time         `facebook:"ts"`
	Extra     core.Params        `facebook:"-"`
}

func (params UpdateURLParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Blacklist != nil {
		out["blacklist"] = *params.Blacklist
	}
	if params.Denylist != nil {
		out["denylist"] = *params.Denylist
	}
	if params.Hmac != nil {
		out["hmac"] = *params.Hmac
	}
	if params.Locale != nil {
		out["locale"] = *params.Locale
	}
	if params.Scopes != nil {
		out["scopes"] = *params.Scopes
	}
	if params.Ts != nil {
		out["ts"] = *params.Ts
	}
	return out
}

func UpdateURL(ctx context.Context, client *core.Client, id string, params UpdateURLParams) (*objects.URL, error) {
	var out objects.URL
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
