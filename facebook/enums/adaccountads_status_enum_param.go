package enums

type AdaccountadsStatusEnumParam string

const (
	AdaccountadsStatusEnumParamActive   AdaccountadsStatusEnumParam = "ACTIVE"
	AdaccountadsStatusEnumParamArchived AdaccountadsStatusEnumParam = "ARCHIVED"
	AdaccountadsStatusEnumParamDeleted  AdaccountadsStatusEnumParam = "DELETED"
	AdaccountadsStatusEnumParamPaused   AdaccountadsStatusEnumParam = "PAUSED"
)

func (value AdaccountadsStatusEnumParam) String() string {
	return string(value)
}
