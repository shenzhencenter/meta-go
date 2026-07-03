package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AnalyticsFunnelQueryResult struct {
	QueryID *core.ID `json:"query_id,omitempty"`
	Status  *string  `json:"status,omitempty"`
}
