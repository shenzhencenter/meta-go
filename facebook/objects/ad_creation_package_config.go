package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreationPackageConfig struct {
	APIVersion                *string  `json:"api_version,omitempty"`
	ID                        *core.ID `json:"id,omitempty"`
	IsEligibleForDefaultOptIn *bool    `json:"is_eligible_for_default_opt_in,omitempty"`
	Objective                 *string  `json:"objective,omitempty"`
	PackageID                 *core.ID `json:"package_id,omitempty"`
	Status                    *string  `json:"status,omitempty"`
}

var AdCreationPackageConfigFields = struct {
	APIVersion                string
	ID                        string
	IsEligibleForDefaultOptIn string
	Objective                 string
	PackageID                 string
	Status                    string
}{
	APIVersion:                "api_version",
	ID:                        "id",
	IsEligibleForDefaultOptIn: "is_eligible_for_default_opt_in",
	Objective:                 "objective",
	PackageID:                 "package_id",
	Status:                    "status",
}
