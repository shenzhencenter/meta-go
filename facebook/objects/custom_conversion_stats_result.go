package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"time"
)

type CustomConversionStatsResult struct {
	Aggregation *enums.CustomconversionstatsresultAggregation `json:"aggregation,omitempty"`
	Data        *[]map[string]interface{}                     `json:"data,omitempty"`
	Timestamp   *time.Time                                    `json:"timestamp,omitempty"`
}
