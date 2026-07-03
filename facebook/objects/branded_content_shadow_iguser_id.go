package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BrandedContentShadowIGUserID struct {
	ID *core.ID `json:"id,omitempty"`
}

var BrandedContentShadowIGUserIDFields = struct {
	ID string
}{
	ID: "id",
}
