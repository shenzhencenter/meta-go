package objects

type AdLimitsEnforcementData struct {
	AdLimitOnPage   *int    `json:"ad_limit_on_page,omitempty"`
	AdLimitOnScope  *int    `json:"ad_limit_on_scope,omitempty"`
	AdVolumeOnPage  *int    `json:"ad_volume_on_page,omitempty"`
	AdVolumeOnScope *int    `json:"ad_volume_on_scope,omitempty"`
	IsAdmin         *bool   `json:"is_admin,omitempty"`
	PageName        *string `json:"page_name,omitempty"`
}

var AdLimitsEnforcementDataFields = struct {
	AdLimitOnPage   string
	AdLimitOnScope  string
	AdVolumeOnPage  string
	AdVolumeOnScope string
	IsAdmin         string
	PageName        string
}{
	AdLimitOnPage:   "ad_limit_on_page",
	AdLimitOnScope:  "ad_limit_on_scope",
	AdVolumeOnPage:  "ad_volume_on_page",
	AdVolumeOnScope: "ad_volume_on_scope",
	IsAdmin:         "is_admin",
	PageName:        "page_name",
}
