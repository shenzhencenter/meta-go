package enums

type EntityattextrangeType string

const (
	EntityattextrangeTypeApplication EntityattextrangeType = "application"
	EntityattextrangeTypeEvent       EntityattextrangeType = "event"
	EntityattextrangeTypeGroup       EntityattextrangeType = "group"
	EntityattextrangeTypePage        EntityattextrangeType = "page"
	EntityattextrangeTypeUser        EntityattextrangeType = "user"
)

func (value EntityattextrangeType) String() string {
	return string(value)
}
