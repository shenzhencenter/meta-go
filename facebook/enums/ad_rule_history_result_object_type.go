package enums

type AdrulehistoryresultObjectType string

const (
	AdrulehistoryresultObjectTypeAd       AdrulehistoryresultObjectType = "AD"
	AdrulehistoryresultObjectTypeAdset    AdrulehistoryresultObjectType = "ADSET"
	AdrulehistoryresultObjectTypeCampaign AdrulehistoryresultObjectType = "CAMPAIGN"
)

func (value AdrulehistoryresultObjectType) String() string {
	return string(value)
}
