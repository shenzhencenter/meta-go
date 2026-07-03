package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AnalyticsFunnelQueryResult struct {
	QueryID *core.ID `json:"query_id,omitempty"`
	Status  *string  `json:"status,omitempty"`
}

var AnalyticsFunnelQueryResultFields = struct {
	QueryID string
	Status  string
}{
	QueryID: "query_id",
	Status:  "status",
}
