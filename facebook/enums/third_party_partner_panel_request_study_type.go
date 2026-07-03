package enums

type ThirdpartypartnerpanelrequestStudyType string

const (
	ThirdpartypartnerpanelrequestStudyTypeBrandLift             ThirdpartypartnerpanelrequestStudyType = "BRAND_LIFT"
	ThirdpartypartnerpanelrequestStudyTypePanelSalesAttribution ThirdpartypartnerpanelrequestStudyType = "PANEL_SALES_ATTRIBUTION"
	ThirdpartypartnerpanelrequestStudyTypeReach                 ThirdpartypartnerpanelrequestStudyType = "REACH"
)

func (value ThirdpartypartnerpanelrequestStudyType) String() string {
	return string(value)
}
