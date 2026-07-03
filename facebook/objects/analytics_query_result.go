package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AnalyticsQueryResult struct {
	QueryID *core.ID `json:"query_id,omitempty"`
	Status  *string  `json:"status,omitempty"`
}
