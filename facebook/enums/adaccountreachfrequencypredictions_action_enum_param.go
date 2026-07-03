package enums

type AdaccountreachfrequencypredictionsActionEnumParam string

const (
	AdaccountreachfrequencypredictionsActionEnumParamCancel  AdaccountreachfrequencypredictionsActionEnumParam = "cancel"
	AdaccountreachfrequencypredictionsActionEnumParamQuote   AdaccountreachfrequencypredictionsActionEnumParam = "quote"
	AdaccountreachfrequencypredictionsActionEnumParamReserve AdaccountreachfrequencypredictionsActionEnumParam = "reserve"
)

func (value AdaccountreachfrequencypredictionsActionEnumParam) String() string {
	return string(value)
}
