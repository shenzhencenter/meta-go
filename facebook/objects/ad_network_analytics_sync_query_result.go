package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdNetworkAnalyticsSyncQueryResult struct {
	OmittedResults *[]map[string]interface{} `json:"omitted_results,omitempty"`
	QueryID        *core.ID                  `json:"query_id,omitempty"`
	Results        *[]map[string]interface{} `json:"results,omitempty"`
}
