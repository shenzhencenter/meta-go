package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PrivateLiftStudyInstance struct {
	BreakdownKey           *string    `json:"breakdown_key,omitempty"`
	CreatedTime            *core.Time `json:"created_time,omitempty"`
	FeatureList            *[]string  `json:"feature_list,omitempty"`
	ID                     *core.ID   `json:"id,omitempty"`
	IssuerCertificate      *string    `json:"issuer_certificate,omitempty"`
	LatestStatusUpdateTime *core.Time `json:"latest_status_update_time,omitempty"`
	RunID                  *core.ID   `json:"run_id,omitempty"`
	ServerHostnames        *[]string  `json:"server_hostnames,omitempty"`
	ServerIps              *[]string  `json:"server_ips,omitempty"`
	Status                 *string    `json:"status,omitempty"`
	Tier                   *string    `json:"tier,omitempty"`
}

var PrivateLiftStudyInstanceFields = struct {
	BreakdownKey           string
	CreatedTime            string
	FeatureList            string
	ID                     string
	IssuerCertificate      string
	LatestStatusUpdateTime string
	RunID                  string
	ServerHostnames        string
	ServerIps              string
	Status                 string
	Tier                   string
}{
	BreakdownKey:           "breakdown_key",
	CreatedTime:            "created_time",
	FeatureList:            "feature_list",
	ID:                     "id",
	IssuerCertificate:      "issuer_certificate",
	LatestStatusUpdateTime: "latest_status_update_time",
	RunID:                  "run_id",
	ServerHostnames:        "server_hostnames",
	ServerIps:              "server_ips",
	Status:                 "status",
	Tier:                   "tier",
}
