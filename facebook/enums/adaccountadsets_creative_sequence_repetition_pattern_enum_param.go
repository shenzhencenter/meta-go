package enums

type AdaccountadsetsCreativeSequenceRepetitionPatternEnumParam string

const (
	AdaccountadsetsCreativeSequenceRepetitionPatternEnumParamFullSequence AdaccountadsetsCreativeSequenceRepetitionPatternEnumParam = "FULL_SEQUENCE"
	AdaccountadsetsCreativeSequenceRepetitionPatternEnumParamLastAd       AdaccountadsetsCreativeSequenceRepetitionPatternEnumParam = "LAST_AD"
)

func (value AdaccountadsetsCreativeSequenceRepetitionPatternEnumParam) String() string {
	return string(value)
}
