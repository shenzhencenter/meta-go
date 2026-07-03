package enums

type ThirdpartypartnerpanelrequestStatus string

const (
	ThirdpartypartnerpanelrequestStatusCreated    ThirdpartypartnerpanelrequestStatus = "CREATED"
	ThirdpartypartnerpanelrequestStatusFailure    ThirdpartypartnerpanelrequestStatus = "FAILURE"
	ThirdpartypartnerpanelrequestStatusInProgress ThirdpartypartnerpanelrequestStatus = "IN_PROGRESS"
	ThirdpartypartnerpanelrequestStatusScheduled  ThirdpartypartnerpanelrequestStatus = "SCHEDULED"
	ThirdpartypartnerpanelrequestStatusSuccess    ThirdpartypartnerpanelrequestStatus = "SUCCESS"
)

func (value ThirdpartypartnerpanelrequestStatus) String() string {
	return string(value)
}
