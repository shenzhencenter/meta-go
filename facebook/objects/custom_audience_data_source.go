package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type CustomAudienceDataSource struct {
	CreationParams *string                                `json:"creation_params,omitempty"`
	SubType        *enums.CustomaudiencedatasourceSubType `json:"sub_type,omitempty"`
	Type           *enums.CustomaudiencedatasourceType    `json:"type,omitempty"`
}
