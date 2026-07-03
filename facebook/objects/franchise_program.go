package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type FranchiseProgram struct {
	BusinessAssetGroup   *BusinessAssetGroup `json:"business_asset_group,omitempty"`
	CreatorBusiness      *Business           `json:"creator_business,omitempty"`
	Description          *string             `json:"description,omitempty"`
	EndDate              *time.Time          `json:"end_date,omitempty"`
	ID                   *core.ID            `json:"id,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	ProgramAccessType    *string             `json:"program_access_type,omitempty"`
	ProgramApprovalType  *string             `json:"program_approval_type,omitempty"`
	ProgramImageLink     *string             `json:"program_image_link,omitempty"`
	ProgramURL           *string             `json:"program_url,omitempty"`
	SharedCustomAudience *CustomAudience     `json:"shared_custom_audience,omitempty"`
	StartDate            *time.Time          `json:"start_date,omitempty"`
}
