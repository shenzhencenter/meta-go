package enums

type AdaccountcampaignsStatusEnumParam string

const (
	AdaccountcampaignsStatusEnumParamActive   AdaccountcampaignsStatusEnumParam = "ACTIVE"
	AdaccountcampaignsStatusEnumParamArchived AdaccountcampaignsStatusEnumParam = "ARCHIVED"
	AdaccountcampaignsStatusEnumParamDeleted  AdaccountcampaignsStatusEnumParam = "DELETED"
	AdaccountcampaignsStatusEnumParamPaused   AdaccountcampaignsStatusEnumParam = "PAUSED"
)

func (value AdaccountcampaignsStatusEnumParam) String() string {
	return string(value)
}
