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

func GetVideoCopyrightUpdateRecordsBatchCall(id string, params GetVideoCopyrightUpdateRecordsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "update_records"), params.ToParams(), options...)
}

func NewGetVideoCopyrightUpdateRecordsBatchRequest(id string, params GetVideoCopyrightUpdateRecordsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MediaCopyrightUpdateRecord]] {
	return core.NewBatchRequest[core.Cursor[objects.MediaCopyrightUpdateRecord]](GetVideoCopyrightUpdateRecordsBatchCall(id, params, options...))
}

func DecodeGetVideoCopyrightUpdateRecordsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MediaCopyrightUpdateRecord], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MediaCopyrightUpdateRecord]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetVideoCopyrightUpdateRecords(ctx context.Context, client *core.Client, id string, params GetVideoCopyrightUpdateRecordsParams) (*core.Cursor[objects.MediaCopyrightUpdateRecord], error) {
	var out core.Cursor[objects.MediaCopyrightUpdateRecord]
	call := GetVideoCopyrightUpdateRecordsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetVideoCopyrightBatchCall(id string, params GetVideoCopyrightParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetVideoCopyrightBatchRequest(id string, params GetVideoCopyrightParams, options ...core.BatchOption) *core.BatchRequest[objects.VideoCopyright] {
	return core.NewBatchRequest[objects.VideoCopyright](GetVideoCopyrightBatchCall(id, params, options...))
}

func DecodeGetVideoCopyrightBatchResponse(response *core.BatchResponse) (*objects.VideoCopyright, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.VideoCopyright
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetVideoCopyright(ctx context.Context, client *core.Client, id string, params GetVideoCopyrightParams) (*objects.VideoCopyright, error) {
	var out objects.VideoCopyright
	call := GetVideoCopyrightBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func UpdateVideoCopyrightBatchCall(id string, params UpdateVideoCopyrightParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateVideoCopyrightBatchRequest(id string, params UpdateVideoCopyrightParams, options ...core.BatchOption) *core.BatchRequest[objects.VideoCopyright] {
	return core.NewBatchRequest[objects.VideoCopyright](UpdateVideoCopyrightBatchCall(id, params, options...))
}

func DecodeUpdateVideoCopyrightBatchResponse(response *core.BatchResponse) (*objects.VideoCopyright, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.VideoCopyright
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateVideoCopyright(ctx context.Context, client *core.Client, id string, params UpdateVideoCopyrightParams) (*objects.VideoCopyright, error) {
	var out objects.VideoCopyright
	call := UpdateVideoCopyrightBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
