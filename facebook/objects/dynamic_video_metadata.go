package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type DynamicVideoMetadata struct {
	ID    *core.ID  `json:"id,omitempty"`
	Tags  *[]string `json:"tags,omitempty"`
	URL   *string   `json:"url,omitempty"`
	Video *AdVideo  `json:"video,omitempty"`
}
