package enums

type JobopeningType string

const (
	JobopeningTypeContract   JobopeningType = "CONTRACT"
	JobopeningTypeFullTime   JobopeningType = "FULL_TIME"
	JobopeningTypeInternship JobopeningType = "INTERNSHIP"
	JobopeningTypePartTime   JobopeningType = "PART_TIME"
	JobopeningTypeVolunteer  JobopeningType = "VOLUNTEER"
)

func (value JobopeningType) String() string {
	return string(value)
}
