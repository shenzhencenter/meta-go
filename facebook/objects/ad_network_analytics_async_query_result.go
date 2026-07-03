package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdNetworkAnalyticsAsyncQueryResult struct {
	Data           *map[string]interface{}   `json:"data,omitempty"`
	OmittedResults *[]map[string]interface{} `json:"omitted_results,omitempty"`
	QueryID        *core.ID                  `json:"query_id,omitempty"`
	Results        *[]map[string]interface{} `json:"results,omitempty"`
	Status         *string                   `json:"status,omitempty"`
}
