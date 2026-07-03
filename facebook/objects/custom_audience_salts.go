package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CustomAudienceSalts struct {
	AppID     *core.ID                  `json:"app_id,omitempty"`
	PublicKey *string                   `json:"public_key,omitempty"`
	Salts     *[]map[string]interface{} `json:"salts,omitempty"`
	UserID    *core.ID                  `json:"user_id,omitempty"`
}
