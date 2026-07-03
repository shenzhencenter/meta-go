package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ConversionActionQuery struct {
	ActionType      *[]map[string]interface{} `json:"action.type,omitempty"`
	Application     *[]map[string]interface{} `json:"application,omitempty"`
	ConversionID    *[]core.ID                `json:"conversion_id,omitempty"`
	Creative        *[]map[string]interface{} `json:"creative,omitempty"`
	Dataset         *[]string                 `json:"dataset,omitempty"`
	DatasetSplit    *[]string                 `json:"dataset_split,omitempty"`
	Event           *[]string                 `json:"event,omitempty"`
	EventCreator    *[]string                 `json:"event.creator,omitempty"`
	EventType       *[]string                 `json:"event_type,omitempty"`
	FbPixel         *[]string                 `json:"fb_pixel,omitempty"`
	FbPixelEvent    *[]string                 `json:"fb_pixel_event,omitempty"`
	Leadgen         *[]string                 `json:"leadgen,omitempty"`
	Object          *[]string                 `json:"object,omitempty"`
	ObjectDomain    *[]string                 `json:"object.domain,omitempty"`
	Offer           *[]string                 `json:"offer,omitempty"`
	OfferCreator    *[]string                 `json:"offer.creator,omitempty"`
	OffsitePixel    *[]string                 `json:"offsite_pixel,omitempty"`
	Page            *[]string                 `json:"page,omitempty"`
	PageParent      *[]string                 `json:"page.parent,omitempty"`
	Post            *[]string                 `json:"post,omitempty"`
	PostObject      *[]string                 `json:"post.object,omitempty"`
	PostObjectWall  *[]string                 `json:"post.object.wall,omitempty"`
	PostWall        *[]string                 `json:"post.wall,omitempty"`
	Question        *[]string                 `json:"question,omitempty"`
	QuestionCreator *[]string                 `json:"question.creator,omitempty"`
	Response        *[]string                 `json:"response,omitempty"`
	Subtype         *[]string                 `json:"subtype,omitempty"`
}
