package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CanvasDynamicSetting struct {
	ChildDocuments *[]Canvas `json:"child_documents,omitempty"`
	ProductSetID   *core.ID  `json:"product_set_id,omitempty"`
}
