package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type PrivateLiftStudyInstance struct {
	BreakdownKey           *string    `json:"breakdown_key,omitempty"`
	CreatedTime            *time.Time `json:"created_time,omitempty"`
	FeatureList            *[]string  `json:"feature_list,omitempty"`
	ID                     *core.ID   `json:"id,omitempty"`
	IssuerCertificate      *string    `json:"issuer_certificate,omitempty"`
	LatestStatusUpdateTime *time.Time `json:"latest_status_update_time,omitempty"`
	RunID                  *core.ID   `json:"run_id,omitempty"`
	ServerHostnames        *[]string  `json:"server_hostnames,omitempty"`
	ServerIps              *[]string  `json:"server_ips,omitempty"`
	Status                 *string    `json:"status,omitempty"`
	Tier                   *string    `json:"tier,omitempty"`
}
