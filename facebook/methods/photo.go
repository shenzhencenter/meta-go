package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetPhotoCommentsParams struct {
	Filter     *enums.PhotocommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.PhotocommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.PhotocommentsOrderEnumParam      `facebook:"order"`
	Since      *time.Time                              `facebook:"since"`
	Extra      core.Params                             `facebook:"-"`
}

func (params GetPhotoCommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	if params.LiveFilter != nil {
		out["live_filter"] = *params.LiveFilter
	}
	if params.Order != nil {
		out["order"] = *params.Order
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	return out
}

func GetPhotoComments(ctx context.Context, client *core.Client, id string, params GetPhotoCommentsParams) (*core.Cursor[objects.Comment], error) {
	var out core.Cursor[objects.Comment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
}

type CreatePhotoCommentsParams struct {
	AttachmentID         *core.ID                                         `facebook:"attachment_id"`
	AttachmentShareURL   *string                                          `facebook:"attachment_share_url"`
	AttachmentURL        *string                                          `facebook:"attachment_url"`
	CommentPrivacyValue  *enums.PhotocommentsCommentPrivacyValueEnumParam `facebook:"comment_privacy_value"`
	FacepileMentionedIds *[]core.ID                                       `facebook:"facepile_mentioned_ids"`
	FeedbackSource       *string                                          `facebook:"feedback_source"`
	IsOffline            *bool                                            `facebook:"is_offline"`
	Message              *string                                          `facebook:"message"`
	NectarModule         *string                                          `facebook:"nectar_module"`
	ObjectID             *core.ID                                         `facebook:"object_id"`
	ParentCommentID      *map[string]interface{}                          `facebook:"parent_comment_id"`
	Text                 *string                                          `facebook:"text"`
	Tracking             *string                                          `facebook:"tracking"`
	Extra                core.Params                                      `facebook:"-"`
}

func (params CreatePhotoCommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AttachmentID != nil {
		out["attachment_id"] = *params.AttachmentID
	}
	if params.AttachmentShareURL != nil {
		out["attachment_share_url"] = *params.AttachmentShareURL
	}
	if params.AttachmentURL != nil {
		out["attachment_url"] = *params.AttachmentURL
	}
	if params.CommentPrivacyValue != nil {
		out["comment_privacy_value"] = *params.CommentPrivacyValue
	}
	if params.FacepileMentionedIds != nil {
		out["facepile_mentioned_ids"] = *params.FacepileMentionedIds
	}
	if params.FeedbackSource != nil {
		out["feedback_source"] = *params.FeedbackSource
	}
	if params.IsOffline != nil {
		out["is_offline"] = *params.IsOffline
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.NectarModule != nil {
		out["nectar_module"] = *params.NectarModule
	}
	if params.ObjectID != nil {
		out["object_id"] = *params.ObjectID
	}
	if params.ParentCommentID != nil {
		out["parent_comment_id"] = *params.ParentCommentID
	}
	if params.Text != nil {
		out["text"] = *params.Text
	}
	if params.Tracking != nil {
		out["tracking"] = *params.Tracking
	}
	return out
}

func CreatePhotoComments(ctx context.Context, client *core.Client, id string, params CreatePhotoCommentsParams) (*objects.Comment, error) {
	var out objects.Comment
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
}

type GetPhotoInsightsParams struct {
	DatePreset *enums.PhotoinsightsDatePresetEnumParam `facebook:"date_preset"`
	Metric     *[]map[string]interface{}               `facebook:"metric"`
	Period     *enums.PhotoinsightsPeriodEnumParam     `facebook:"period"`
	Since      *time.Time                              `facebook:"since"`
	Until      *time.Time                              `facebook:"until"`
	Extra      core.Params                             `facebook:"-"`
}

func (params GetPhotoInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DatePreset != nil {
		out["date_preset"] = *params.DatePreset
	}
	if params.Metric != nil {
		out["metric"] = *params.Metric
	}
	if params.Period != nil {
		out["period"] = *params.Period
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetPhotoInsights(ctx context.Context, client *core.Client, id string, params GetPhotoInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	var out core.Cursor[objects.InsightsResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), &out)
	return &out, err
}

type GetPhotoLikesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPhotoLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPhotoLikes(ctx context.Context, client *core.Client, id string, params GetPhotoLikesParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
}

type CreatePhotoLikesParams struct {
	AttributionIDV2 *string     `facebook:"attribution_id_v2"`
	FeedbackSource  *string     `facebook:"feedback_source"`
	NectarModule    *string     `facebook:"nectar_module"`
	Notify          *bool       `facebook:"notify"`
	Tracking        *string     `facebook:"tracking"`
	Extra           core.Params `facebook:"-"`
}

func (params CreatePhotoLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AttributionIDV2 != nil {
		out["attribution_id_v2"] = *params.AttributionIDV2
	}
	if params.FeedbackSource != nil {
		out["feedback_source"] = *params.FeedbackSource
	}
	if params.NectarModule != nil {
		out["nectar_module"] = *params.NectarModule
	}
	if params.Notify != nil {
		out["notify"] = *params.Notify
	}
	if params.Tracking != nil {
		out["tracking"] = *params.Tracking
	}
	return out
}

func CreatePhotoLikes(ctx context.Context, client *core.Client, id string, params CreatePhotoLikesParams) (*objects.Photo, error) {
	var out objects.Photo
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
}

type GetPhotoSponsorTagsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPhotoSponsorTagsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPhotoSponsorTags(ctx context.Context, client *core.Client, id string, params GetPhotoSponsorTagsParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "sponsor_tags"), params.ToParams(), &out)
	return &out, err
}

type DeletePhotoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeletePhotoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeletePhoto(ctx context.Context, client *core.Client, id string, params DeletePhotoParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetPhotoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPhotoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPhoto(ctx context.Context, client *core.Client, id string, params GetPhotoParams) (*objects.Photo, error) {
	var out objects.Photo
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
