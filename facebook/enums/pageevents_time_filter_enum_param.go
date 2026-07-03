package enums

type PageeventsTimeFilterEnumParam string

const (
	PageeventsTimeFilterEnumParamPast     PageeventsTimeFilterEnumParam = "past"
	PageeventsTimeFilterEnumParamUpcoming PageeventsTimeFilterEnumParam = "upcoming"
)

func (value PageeventsTimeFilterEnumParam) String() string {
	return string(value)
}
