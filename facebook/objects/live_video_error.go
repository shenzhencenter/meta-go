package objects

import (
	"time"
)

type LiveVideoError struct {
	CreationTime *time.Time `json:"creation_time,omitempty"`
	ErrorCode    *int       `json:"error_code,omitempty"`
	ErrorMessage *string    `json:"error_message,omitempty"`
	ErrorType    *string    `json:"error_type,omitempty"`
}
