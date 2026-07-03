package enums

type AdcampaignCreativeSequenceRepetitionPattern string

const (
	AdcampaignCreativeSequenceRepetitionPatternFullSequence AdcampaignCreativeSequenceRepetitionPattern = "FULL_SEQUENCE"
	AdcampaignCreativeSequenceRepetitionPatternLastAd       AdcampaignCreativeSequenceRepetitionPattern = "LAST_AD"
)

func (value AdcampaignCreativeSequenceRepetitionPattern) String() string {
	return string(value)
}
