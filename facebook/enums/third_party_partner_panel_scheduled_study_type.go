package enums

type ThirdpartypartnerpanelscheduledStudyType string

const (
	ThirdpartypartnerpanelscheduledStudyTypeBrandLift             ThirdpartypartnerpanelscheduledStudyType = "BRAND_LIFT"
	ThirdpartypartnerpanelscheduledStudyTypePanelSalesAttribution ThirdpartypartnerpanelscheduledStudyType = "PANEL_SALES_ATTRIBUTION"
	ThirdpartypartnerpanelscheduledStudyTypeReach                 ThirdpartypartnerpanelscheduledStudyType = "REACH"
)

func (value ThirdpartypartnerpanelscheduledStudyType) String() string {
	return string(value)
}
