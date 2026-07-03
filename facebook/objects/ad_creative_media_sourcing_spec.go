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
