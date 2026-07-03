package enums

type JobopeningJobStatus string

const (
	JobopeningJobStatusClosed      JobopeningJobStatus = "CLOSED"
	JobopeningJobStatusDraft       JobopeningJobStatus = "DRAFT"
	JobopeningJobStatusOpen        JobopeningJobStatus = "OPEN"
	JobopeningJobStatusProvisional JobopeningJobStatus = "PROVISIONAL"
)

func (value JobopeningJobStatus) String() string {
	return string(value)
}
