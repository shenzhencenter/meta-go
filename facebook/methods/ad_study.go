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

func GetAdStudyCells(ctx context.Context, client *core.Client, id string, params GetAdStudyCellsParams) (*core.Cursor[objects.AdStudyCell], error) {
	var out core.Cursor[objects.AdStudyCell]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "cells"), params.ToParams(), &out)
	return &out, err
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

func CreateAdStudyCheckpoint(ctx context.Context, client *core.Client, id string, params CreateAdStudyCheckpointParams) (*objects.AdStudy, error) {
	var out objects.AdStudy
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "checkpoint"), params.ToParams(), &out)
	return &out, err
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

func GetAdStudyInstances(ctx context.Context, client *core.Client, id string, params GetAdStudyInstancesParams) (*core.Cursor[objects.PrivateLiftStudyInstance], error) {
	var out core.Cursor[objects.PrivateLiftStudyInstance]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "instances"), params.ToParams(), &out)
	return &out, err
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

func CreateAdStudyInstances(ctx context.Context, client *core.Client, id string, params CreateAdStudyInstancesParams) (*objects.PrivateLiftStudyInstance, error) {
	var out objects.PrivateLiftStudyInstance
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "instances"), params.ToParams(), &out)
	return &out, err
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

func GetAdStudyObjectives(ctx context.Context, client *core.Client, id string, params GetAdStudyObjectivesParams) (*core.Cursor[objects.AdStudyObjective], error) {
	var out core.Cursor[objects.AdStudyObjective]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "objectives"), params.ToParams(), &out)
	return &out, err
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

func DeleteAdStudy(ctx context.Context, client *core.Client, id string, params DeleteAdStudyParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func GetAdStudy(ctx context.Context, client *core.Client, id string, params GetAdStudyParams) (*objects.AdStudy, error) {
	var out objects.AdStudy
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func UpdateAdStudy(ctx context.Context, client *core.Client, id string, params UpdateAdStudyParams) (*objects.AdStudy, error) {
	var out objects.AdStudy
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
