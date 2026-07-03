package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CustomAudienceAdAccount struct {
	ID *core.ID `json:"id,omitempty"`
}

var CustomAudienceAdAccountFields = struct {
	ID string
}{
	ID: "id",
}
