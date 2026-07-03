package objects

type AdLimitsEnforcementData struct {
	AdLimitOnPage   *int    `json:"ad_limit_on_page,omitempty"`
	AdLimitOnScope  *int    `json:"ad_limit_on_scope,omitempty"`
	AdVolumeOnPage  *int    `json:"ad_volume_on_page,omitempty"`
	AdVolumeOnScope *int    `json:"ad_volume_on_scope,omitempty"`
	IsAdmin         *bool   `json:"is_admin,omitempty"`
	PageName        *string `json:"page_name,omitempty"`
}
