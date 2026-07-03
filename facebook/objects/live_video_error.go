package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LiveVideoError struct {
	CreationTime *core.Time `json:"creation_time,omitempty"`
	ErrorCode    *int       `json:"error_code,omitempty"`
	ErrorMessage *string    `json:"error_message,omitempty"`
	ErrorType    *string    `json:"error_type,omitempty"`
}

var LiveVideoErrorFields = struct {
	CreationTime string
	ErrorCode    string
	ErrorMessage string
	ErrorType    string
}{
	CreationTime: "creation_time",
	ErrorCode:    "error_code",
	ErrorMessage: "error_message",
	ErrorType:    "error_type",
}
