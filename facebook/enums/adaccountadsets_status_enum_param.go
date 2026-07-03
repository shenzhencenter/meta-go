package enums

type AdaccountadsetsStatusEnumParam string

const (
	AdaccountadsetsStatusEnumParamActive   AdaccountadsetsStatusEnumParam = "ACTIVE"
	AdaccountadsetsStatusEnumParamArchived AdaccountadsetsStatusEnumParam = "ARCHIVED"
	AdaccountadsetsStatusEnumParamDeleted  AdaccountadsetsStatusEnumParam = "DELETED"
	AdaccountadsetsStatusEnumParamPaused   AdaccountadsetsStatusEnumParam = "PAUSED"
)

func (value AdaccountadsetsStatusEnumParam) String() string {
	return string(value)
}
