package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CustomAudienceSalts struct {
	AppID     *core.ID                  `json:"app_id,omitempty"`
	PublicKey *string                   `json:"public_key,omitempty"`
	Salts     *[]map[string]interface{} `json:"salts,omitempty"`
	UserID    *core.ID                  `json:"user_id,omitempty"`
}

var CustomAudienceSaltsFields = struct {
	AppID     string
	PublicKey string
	Salts     string
	UserID    string
}{
	AppID:     "app_id",
	PublicKey: "public_key",
	Salts:     "salts",
	UserID:    "user_id",
}
