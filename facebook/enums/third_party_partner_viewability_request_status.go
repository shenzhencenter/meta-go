package enums

type ThirdpartypartnerviewabilityrequestStatus string

const (
	ThirdpartypartnerviewabilityrequestStatusCreated    ThirdpartypartnerviewabilityrequestStatus = "CREATED"
	ThirdpartypartnerviewabilityrequestStatusFailure    ThirdpartypartnerviewabilityrequestStatus = "FAILURE"
	ThirdpartypartnerviewabilityrequestStatusInProgress ThirdpartypartnerviewabilityrequestStatus = "IN_PROGRESS"
	ThirdpartypartnerviewabilityrequestStatusScheduled  ThirdpartypartnerviewabilityrequestStatus = "SCHEDULED"
	ThirdpartypartnerviewabilityrequestStatusSuccess    ThirdpartypartnerviewabilityrequestStatus = "SUCCESS"
)

func (value ThirdpartypartnerviewabilityrequestStatus) String() string {
	return string(value)
}
