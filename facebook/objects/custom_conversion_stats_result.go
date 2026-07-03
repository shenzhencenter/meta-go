package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type CustomConversionStatsResult struct {
	Aggregation *enums.CustomconversionstatsresultAggregation `json:"aggregation,omitempty"`
	Data        *[]map[string]interface{}                     `json:"data,omitempty"`
	Timestamp   *core.Time                                    `json:"timestamp,omitempty"`
}

var CustomConversionStatsResultFields = struct {
	Aggregation string
	Data        string
	Timestamp   string
}{
	Aggregation: "aggregation",
	Data:        "data",
	Timestamp:   "timestamp",
}
