package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type HasLeadAccess struct {
	AppHasLeadsPermission    *bool    `json:"app_has_leads_permission,omitempty"`
	CanAccessLead            *bool    `json:"can_access_lead,omitempty"`
	EnabledLeadAccessManager *bool    `json:"enabled_lead_access_manager,omitempty"`
	FailureReason            *string  `json:"failure_reason,omitempty"`
	FailureResolution        *string  `json:"failure_resolution,omitempty"`
	IsPageAdmin              *bool    `json:"is_page_admin,omitempty"`
	PageID                   *core.ID `json:"page_id,omitempty"`
	UserHasLeadsPermission   *bool    `json:"user_has_leads_permission,omitempty"`
	UserID                   *core.ID `json:"user_id,omitempty"`
}
