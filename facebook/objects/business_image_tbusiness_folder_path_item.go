package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessImageTBusinessFolderPathItem struct {
	ID             *core.ID `json:"id,omitempty"`
	ParentFolderID *core.ID `json:"parent_folder_id,omitempty"`
	Type           *string  `json:"type,omitempty"`
}
