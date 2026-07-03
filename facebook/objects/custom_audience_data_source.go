package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type CustomAudienceDataSource struct {
	CreationParams *string                                `json:"creation_params,omitempty"`
	SubType        *enums.CustomaudiencedatasourceSubType `json:"sub_type,omitempty"`
	Type           *enums.CustomaudiencedatasourceType    `json:"type,omitempty"`
}

var CustomAudienceDataSourceFields = struct {
	CreationParams string
	SubType        string
	Type           string
}{
	CreationParams: "creation_params",
	SubType:        "sub_type",
	Type:           "type",
}
