package enums

type PagemessagesSuggestionActionEnumParam string

const (
	PagemessagesSuggestionActionEnumParamAccept     PagemessagesSuggestionActionEnumParam = "ACCEPT"
	PagemessagesSuggestionActionEnumParamDismiss    PagemessagesSuggestionActionEnumParam = "DISMISS"
	PagemessagesSuggestionActionEnumParamImpression PagemessagesSuggestionActionEnumParam = "IMPRESSION"
)

func (value PagemessagesSuggestionActionEnumParam) String() string {
	return string(value)
}
