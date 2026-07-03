package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AnalyticsUserConfig struct {
	DemoAppNuxConfig *map[string]interface{} `json:"demo_app_nux_config,omitempty"`
	Flags            *[]map[string]string    `json:"flags,omitempty"`
	ID               *core.ID                `json:"id,omitempty"`
}
