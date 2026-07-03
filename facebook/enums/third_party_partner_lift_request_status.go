package enums

type ThirdpartypartnerliftrequestStatus string

const (
	ThirdpartypartnerliftrequestStatusCreated    ThirdpartypartnerliftrequestStatus = "CREATED"
	ThirdpartypartnerliftrequestStatusFailure    ThirdpartypartnerliftrequestStatus = "FAILURE"
	ThirdpartypartnerliftrequestStatusInProgress ThirdpartypartnerliftrequestStatus = "IN_PROGRESS"
	ThirdpartypartnerliftrequestStatusScheduled  ThirdpartypartnerliftrequestStatus = "SCHEDULED"
	ThirdpartypartnerliftrequestStatusSuccess    ThirdpartypartnerliftrequestStatus = "SUCCESS"
)

func (value ThirdpartypartnerliftrequestStatus) String() string {
	return string(value)
}
