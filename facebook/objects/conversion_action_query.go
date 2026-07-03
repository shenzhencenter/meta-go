package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var ConversionActionQueryFields = struct {
	ActionType      string
	Application     string
	ConversionID    string
	Creative        string
	Dataset         string
	DatasetSplit    string
	Event           string
	EventCreator    string
	EventType       string
	FbPixel         string
	FbPixelEvent    string
	Leadgen         string
	Object          string
	ObjectDomain    string
	Offer           string
	OfferCreator    string
	OffsitePixel    string
	Page            string
	PageParent      string
	Post            string
	PostObject      string
	PostObjectWall  string
	PostWall        string
	Question        string
	QuestionCreator string
	Response        string
	Subtype         string
}{
	ActionType:      "action.type",
	Application:     "application",
	ConversionID:    "conversion_id",
	Creative:        "creative",
	Dataset:         "dataset",
	DatasetSplit:    "dataset_split",
	Event:           "event",
	EventCreator:    "event.creator",
	EventType:       "event_type",
	FbPixel:         "fb_pixel",
	FbPixelEvent:    "fb_pixel_event",
	Leadgen:         "leadgen",
	Object:          "object",
	ObjectDomain:    "object.domain",
	Offer:           "offer",
	OfferCreator:    "offer.creator",
	OffsitePixel:    "offsite_pixel",
	Page:            "page",
	PageParent:      "page.parent",
	Post:            "post",
	PostObject:      "post.object",
	PostObjectWall:  "post.object.wall",
	PostWall:        "post.wall",
	Question:        "question",
	QuestionCreator: "question.creator",
	Response:        "response",
	Subtype:         "subtype",
}
