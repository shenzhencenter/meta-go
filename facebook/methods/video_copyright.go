package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetVideoCopyrightUpdateRecordsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVideoCopyrightUpdateRecordsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVideoCopyrightUpdateRecords(ctx context.Context, client *core.Client, id string, params GetVideoCopyrightUpdateRecordsParams) (*core.Cursor[objects.MediaCopyrightUpdateRecord], error) {
	var out core.Cursor[objects.MediaCopyrightUpdateRecord]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "update_records"), params.ToParams(), &out)
	return &out, err
}

type GetVideoCopyrightParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVideoCopyrightParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVideoCopyright(ctx context.Context, client *core.Client, id string, params GetVideoCopyrightParams) (*objects.VideoCopyright, error) {
	var out objects.VideoCopyright
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateVideoCopyrightParams struct {
	AppendExcludedOwnershipSegments *bool                                `facebook:"append_excluded_ownership_segments"`
	AttributionID                   *core.ID                             `facebook:"attribution_id"`
	ContentCategory                 *enums.VideocopyrightContentCategory `facebook:"content_category"`
	ExcludedOwnershipCountries      *[]string                            `facebook:"excluded_ownership_countries"`
	ExcludedOwnershipSegments       *[]map[string]interface{}            `facebook:"excluded_ownership_segments"`
	IsReferenceDisabled             *bool                                `facebook:"is_reference_disabled"`
	MonitoringType                  *enums.VideocopyrightMonitoringType  `facebook:"monitoring_type"`
	OwnershipCountries              *[]string                            `facebook:"ownership_countries"`
	RuleID                          *core.ID                             `facebook:"rule_id"`
	WhitelistedIds                  *[]core.ID                           `facebook:"whitelisted_ids"`
	WhitelistedIgUserIds            *[]core.ID                           `facebook:"whitelisted_ig_user_ids"`
	Extra                           core.Params                          `facebook:"-"`
}

func (params UpdateVideoCopyrightParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AppendExcludedOwnershipSegments != nil {
		out["append_excluded_ownership_segments"] = *params.AppendExcludedOwnershipSegments
	}
	if params.AttributionID != nil {
		out["attribution_id"] = *params.AttributionID
	}
	if params.ContentCategory != nil {
		out["content_category"] = *params.ContentCategory
	}
	if params.ExcludedOwnershipCountries != nil {
		out["excluded_ownership_countries"] = *params.ExcludedOwnershipCountries
	}
	if params.ExcludedOwnershipSegments != nil {
		out["excluded_ownership_segments"] = *params.ExcludedOwnershipSegments
	}
	if params.IsReferenceDisabled != nil {
		out["is_reference_disabled"] = *params.IsReferenceDisabled
	}
	if params.MonitoringType != nil {
		out["monitoring_type"] = *params.MonitoringType
	}
	if params.OwnershipCountries != nil {
		out["ownership_countries"] = *params.OwnershipCountries
	}
	if params.RuleID != nil {
		out["rule_id"] = *params.RuleID
	}
	if params.WhitelistedIds != nil {
		out["whitelisted_ids"] = *params.WhitelistedIds
	}
	if params.WhitelistedIgUserIds != nil {
		out["whitelisted_ig_user_ids"] = *params.WhitelistedIgUserIds
	}
	return out
}

func UpdateVideoCopyright(ctx context.Context, client *core.Client, id string, params UpdateVideoCopyrightParams) (*objects.VideoCopyright, error) {
	var out objects.VideoCopyright
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
