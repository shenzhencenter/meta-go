package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeMediaSourcingSpec struct {
	Bodies          *[]map[string]interface{} `json:"bodies,omitempty"`
	Descriptions    *[]map[string]interface{} `json:"descriptions,omitempty"`
	Images          *[]map[string]interface{} `json:"images,omitempty"`
	PushMetadataIds *[]core.ID                `json:"push_metadata_ids,omitempty"`
	RelatedMedia    *map[string]interface{}   `json:"related_media,omitempty"`
	Titles          *[]map[string]interface{} `json:"titles,omitempty"`
	Videos          *[]map[string]interface{} `json:"videos,omitempty"`
}

var AdCreativeMediaSourcingSpecFields = struct {
	Bodies          string
	Descriptions    string
	Images          string
	PushMetadataIds string
	RelatedMedia    string
	Titles          string
	Videos          string
}{
	Bodies:          "bodies",
	Descriptions:    "descriptions",
	Images:          "images",
	PushMetadataIds: "push_metadata_ids",
	RelatedMedia:    "related_media",
	Titles:          "titles",
	Videos:          "videos",
}
