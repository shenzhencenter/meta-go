package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductFeedUploadProgress struct {
	Pos         *int       `json:"pos,omitempty"`
	Size        *int       `json:"size,omitempty"`
	Step        *string    `json:"step,omitempty"`
	Unit        *string    `json:"unit,omitempty"`
	UpdatedTime *core.Time `json:"updated_time,omitempty"`
}

var ProductFeedUploadProgressFields = struct {
	Pos         string
	Size        string
	Step        string
	Unit        string
	UpdatedTime string
}{
	Pos:         "pos",
	Size:        "size",
	Step:        "step",
	Unit:        "unit",
	UpdatedTime: "updated_time",
}
