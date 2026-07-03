package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type FranchiseProgram struct {
	BusinessAssetGroup   *BusinessAssetGroup `json:"business_asset_group,omitempty"`
	CreatorBusiness      *Business           `json:"creator_business,omitempty"`
	Description          *string             `json:"description,omitempty"`
	EndDate              *core.Time          `json:"end_date,omitempty"`
	ID                   *core.ID            `json:"id,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	ProgramAccessType    *string             `json:"program_access_type,omitempty"`
	ProgramApprovalType  *string             `json:"program_approval_type,omitempty"`
	ProgramImageLink     *string             `json:"program_image_link,omitempty"`
	ProgramURL           *string             `json:"program_url,omitempty"`
	SharedCustomAudience *CustomAudience     `json:"shared_custom_audience,omitempty"`
	StartDate            *core.Time          `json:"start_date,omitempty"`
}

var FranchiseProgramFields = struct {
	BusinessAssetGroup   string
	CreatorBusiness      string
	Description          string
	EndDate              string
	ID                   string
	Name                 string
	ProgramAccessType    string
	ProgramApprovalType  string
	ProgramImageLink     string
	ProgramURL           string
	SharedCustomAudience string
	StartDate            string
}{
	BusinessAssetGroup:   "business_asset_group",
	CreatorBusiness:      "creator_business",
	Description:          "description",
	EndDate:              "end_date",
	ID:                   "id",
	Name:                 "name",
	ProgramAccessType:    "program_access_type",
	ProgramApprovalType:  "program_approval_type",
	ProgramImageLink:     "program_image_link",
	ProgramURL:           "program_url",
	SharedCustomAudience: "shared_custom_audience",
	StartDate:            "start_date",
}
