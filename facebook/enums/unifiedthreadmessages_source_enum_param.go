package enums

type UnifiedthreadmessagesSourceEnumParam string

const (
	UnifiedthreadmessagesSourceEnumParamAll          UnifiedthreadmessagesSourceEnumParam = "ALL"
	UnifiedthreadmessagesSourceEnumParamParticipants UnifiedthreadmessagesSourceEnumParam = "PARTICIPANTS"
)

func (value UnifiedthreadmessagesSourceEnumParam) String() string {
	return string(value)
}
