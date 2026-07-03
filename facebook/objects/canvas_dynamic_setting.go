package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CanvasDynamicSetting struct {
	ChildDocuments *[]Canvas `json:"child_documents,omitempty"`
	ProductSetID   *core.ID  `json:"product_set_id,omitempty"`
}

var CanvasDynamicSettingFields = struct {
	ChildDocuments string
	ProductSetID   string
}{
	ChildDocuments: "child_documents",
	ProductSetID:   "product_set_id",
}
