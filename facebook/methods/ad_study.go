package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdStudyCellsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyCellsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyCellsBatchCall(id string, params GetAdStudyCellsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "cells"), params.ToParams(), options...)
}

func NewGetAdStudyCellsBatchRequest(id string, params GetAdStudyCellsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdStudyCell]] {
	return core.NewBatchRequest[core.Cursor[objects.AdStudyCell]](GetAdStudyCellsBatchCall(id, params, options...))
}

func DecodeGetAdStudyCellsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdStudyCell], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdStudyCell]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyCellsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdStudyCellsParams) (*core.Cursor[objects.AdStudyCell], *core.Response, error) {
	var out core.Cursor[objects.AdStudyCell]
	call := GetAdStudyCellsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdStudyCells(ctx context.Context, client *core.Client, id string, params GetAdStudyCellsParams) (*core.Cursor[objects.AdStudyCell], error) {
	out, _, err := GetAdStudyCellsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdStudyCheckpointParams struct {
	CheckpointData string      `facebook:"checkpoint_data"`
	CheckpointName string      `facebook:"checkpoint_name"`
	Component      string      `facebook:"component"`
	InstanceID     *core.ID    `facebook:"instance_id"`
	RunID          *core.ID    `facebook:"run_id"`
	Extra          core.Params `facebook:"-"`
}

func (params CreateAdStudyCheckpointParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["checkpoint_data"] = params.CheckpointData
	out["checkpoint_name"] = params.CheckpointName
	out["component"] = params.Component
	if params.InstanceID != nil {
		out["instance_id"] = *params.InstanceID
	}
	if params.RunID != nil {
		out["run_id"] = *params.RunID
	}
	return out
}

func CreateAdStudyCheckpointBatchCall(id string, params CreateAdStudyCheckpointParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "checkpoint"), params.ToParams(), options...)
}

func NewCreateAdStudyCheckpointBatchRequest(id string, params CreateAdStudyCheckpointParams, options ...core.BatchOption) *core.BatchRequest[objects.AdStudy] {
	return core.NewBatchRequest[objects.AdStudy](CreateAdStudyCheckpointBatchCall(id, params, options...))
}

func DecodeCreateAdStudyCheckpointBatchResponse(response *core.BatchResponse) (*objects.AdStudy, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdStudy
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdStudyCheckpointWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdStudyCheckpointParams) (*objects.AdStudy, *core.Response, error) {
	var out objects.AdStudy
	call := CreateAdStudyCheckpointBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdStudyCheckpoint(ctx context.Context, client *core.Client, id string, params CreateAdStudyCheckpointParams) (*objects.AdStudy, error) {
	out, _, err := CreateAdStudyCheckpointWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdStudyInstancesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyInstancesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyInstancesBatchCall(id string, params GetAdStudyInstancesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "instances"), params.ToParams(), options...)
}

func NewGetAdStudyInstancesBatchRequest(id string, params GetAdStudyInstancesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PrivateLiftStudyInstance]] {
	return core.NewBatchRequest[core.Cursor[objects.PrivateLiftStudyInstance]](GetAdStudyInstancesBatchCall(id, params, options...))
}

func DecodeGetAdStudyInstancesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PrivateLiftStudyInstance], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PrivateLiftStudyInstance]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyInstancesWithResponse(ctx context.Context, client *core.Client, id string, params GetAdStudyInstancesParams) (*core.Cursor[objects.PrivateLiftStudyInstance], *core.Response, error) {
	var out core.Cursor[objects.PrivateLiftStudyInstance]
	call := GetAdStudyInstancesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdStudyInstances(ctx context.Context, client *core.Client, id string, params GetAdStudyInstancesParams) (*core.Cursor[objects.PrivateLiftStudyInstance], error) {
	out, _, err := GetAdStudyInstancesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdStudyInstancesParams struct {
	BreakdownKey map[string]interface{} `facebook:"breakdown_key"`
	RunID        *core.ID               `facebook:"run_id"`
	Extra        core.Params            `facebook:"-"`
}

func (params CreateAdStudyInstancesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["breakdown_key"] = params.BreakdownKey
	if params.RunID != nil {
		out["run_id"] = *params.RunID
	}
	return out
}

func CreateAdStudyInstancesBatchCall(id string, params CreateAdStudyInstancesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "instances"), params.ToParams(), options...)
}

func NewCreateAdStudyInstancesBatchRequest(id string, params CreateAdStudyInstancesParams, options ...core.BatchOption) *core.BatchRequest[objects.PrivateLiftStudyInstance] {
	return core.NewBatchRequest[objects.PrivateLiftStudyInstance](CreateAdStudyInstancesBatchCall(id, params, options...))
}

func DecodeCreateAdStudyInstancesBatchResponse(response *core.BatchResponse) (*objects.PrivateLiftStudyInstance, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PrivateLiftStudyInstance
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdStudyInstancesWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdStudyInstancesParams) (*objects.PrivateLiftStudyInstance, *core.Response, error) {
	var out objects.PrivateLiftStudyInstance
	call := CreateAdStudyInstancesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdStudyInstances(ctx context.Context, client *core.Client, id string, params CreateAdStudyInstancesParams) (*objects.PrivateLiftStudyInstance, error) {
	out, _, err := CreateAdStudyInstancesWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdStudyObjectivesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyObjectivesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyObjectivesBatchCall(id string, params GetAdStudyObjectivesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "objectives"), params.ToParams(), options...)
}

func NewGetAdStudyObjectivesBatchRequest(id string, params GetAdStudyObjectivesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdStudyObjective]] {
	return core.NewBatchRequest[core.Cursor[objects.AdStudyObjective]](GetAdStudyObjectivesBatchCall(id, params, options...))
}

func DecodeGetAdStudyObjectivesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdStudyObjective], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdStudyObjective]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyObjectivesWithResponse(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectivesParams) (*core.Cursor[objects.AdStudyObjective], *core.Response, error) {
	var out core.Cursor[objects.AdStudyObjective]
	call := GetAdStudyObjectivesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdStudyObjectives(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectivesParams) (*core.Cursor[objects.AdStudyObjective], error) {
	out, _, err := GetAdStudyObjectivesWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteAdStudyParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdStudyParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdStudyBatchCall(id string, params DeleteAdStudyParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteAdStudyBatchRequest(id string, params DeleteAdStudyParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdStudyBatchCall(id, params, options...))
}

func DecodeDeleteAdStudyBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteAdStudyWithResponse(ctx context.Context, client *core.Client, id string, params DeleteAdStudyParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteAdStudyBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteAdStudy(ctx context.Context, client *core.Client, id string, params DeleteAdStudyParams) (*map[string]interface{}, error) {
	out, _, err := DeleteAdStudyWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdStudyParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyBatchCall(id string, params GetAdStudyParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdStudyBatchRequest(id string, params GetAdStudyParams, options ...core.BatchOption) *core.BatchRequest[objects.AdStudy] {
	return core.NewBatchRequest[objects.AdStudy](GetAdStudyBatchCall(id, params, options...))
}

func DecodeGetAdStudyBatchResponse(response *core.BatchResponse) (*objects.AdStudy, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdStudy
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyWithResponse(ctx context.Context, client *core.Client, id string, params GetAdStudyParams) (*objects.AdStudy, *core.Response, error) {
	var out objects.AdStudy
	call := GetAdStudyBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdStudy(ctx context.Context, client *core.Client, id string, params GetAdStudyParams) (*objects.AdStudy, error) {
	out, _, err := GetAdStudyWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateAdStudyParams struct {
	Cells              *[]map[string]interface{} `facebook:"cells"`
	ClientBusiness     *string                   `facebook:"client_business"`
	ConfidenceLevel    *float64                  `facebook:"confidence_level"`
	CooldownStartTime  *int                      `facebook:"cooldown_start_time"`
	CreativeTestConfig *map[string]interface{}   `facebook:"creative_test_config"`
	Description        *string                   `facebook:"description"`
	EndTime            *int                      `facebook:"end_time"`
	Name               *string                   `facebook:"name"`
	Objectives         *[]map[string]interface{} `facebook:"objectives"`
	ObservationEndTime *int                      `facebook:"observation_end_time"`
	StartTime          *int                      `facebook:"start_time"`
	Type               *enums.AdstudyType        `facebook:"type"`
	Viewers            *[]int                    `facebook:"viewers"`
	Extra              core.Params               `facebook:"-"`
}

func (params UpdateAdStudyParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Cells != nil {
		out["cells"] = *params.Cells
	}
	if params.ClientBusiness != nil {
		out["client_business"] = *params.ClientBusiness
	}
	if params.ConfidenceLevel != nil {
		out["confidence_level"] = *params.ConfidenceLevel
	}
	if params.CooldownStartTime != nil {
		out["cooldown_start_time"] = *params.CooldownStartTime
	}
	if params.CreativeTestConfig != nil {
		out["creative_test_config"] = *params.CreativeTestConfig
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.Objectives != nil {
		out["objectives"] = *params.Objectives
	}
	if params.ObservationEndTime != nil {
		out["observation_end_time"] = *params.ObservationEndTime
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	if params.Viewers != nil {
		out["viewers"] = *params.Viewers
	}
	return out
}

func UpdateAdStudyBatchCall(id string, params UpdateAdStudyParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateAdStudyBatchRequest(id string, params UpdateAdStudyParams, options ...core.BatchOption) *core.BatchRequest[objects.AdStudy] {
	return core.NewBatchRequest[objects.AdStudy](UpdateAdStudyBatchCall(id, params, options...))
}

func DecodeUpdateAdStudyBatchResponse(response *core.BatchResponse) (*objects.AdStudy, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdStudy
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateAdStudyWithResponse(ctx context.Context, client *core.Client, id string, params UpdateAdStudyParams) (*objects.AdStudy, *core.Response, error) {
	var out objects.AdStudy
	call := UpdateAdStudyBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateAdStudy(ctx context.Context, client *core.Client, id string, params UpdateAdStudyParams) (*objects.AdStudy, error) {
	out, _, err := UpdateAdStudyWithResponse(ctx, client, id, params)
	return out, err
}
