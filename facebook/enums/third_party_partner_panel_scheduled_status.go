package enums

type ThirdpartypartnerpanelscheduledStatus string

const (
	ThirdpartypartnerpanelscheduledStatusCancelled ThirdpartypartnerpanelscheduledStatus = "CANCELLED"
	ThirdpartypartnerpanelscheduledStatusCreated   ThirdpartypartnerpanelscheduledStatus = "CREATED"
	ThirdpartypartnerpanelscheduledStatusFinished  ThirdpartypartnerpanelscheduledStatus = "FINISHED"
	ThirdpartypartnerpanelscheduledStatusOngoing   ThirdpartypartnerpanelscheduledStatus = "ONGOING"
)

func (value ThirdpartypartnerpanelscheduledStatus) String() string {
	return string(value)
}
