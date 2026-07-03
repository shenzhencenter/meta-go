package objects

import (
	"time"
)

type ProductFeedUploadProgress struct {
	Pos         *int       `json:"pos,omitempty"`
	Size        *int       `json:"size,omitempty"`
	Step        *string    `json:"step,omitempty"`
	Unit        *string    `json:"unit,omitempty"`
	UpdatedTime *time.Time `json:"updated_time,omitempty"`
}
