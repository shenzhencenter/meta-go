package enums

type AdcampaignplacementgetEffectiveMessengerPositions string

const (
	AdcampaignplacementgetEffectiveMessengerPositionsMessengerInbox             AdcampaignplacementgetEffectiveMessengerPositions = "MESSENGER_INBOX"
	AdcampaignplacementgetEffectiveMessengerPositionsMessengerMarketingMessages AdcampaignplacementgetEffectiveMessengerPositions = "MESSENGER_MARKETING_MESSAGES"
	AdcampaignplacementgetEffectiveMessengerPositionsMessengerStory             AdcampaignplacementgetEffectiveMessengerPositions = "MESSENGER_STORY"
	AdcampaignplacementgetEffectiveMessengerPositionsMessengerThread            AdcampaignplacementgetEffectiveMessengerPositions = "MESSENGER_THREAD"
)

func (value AdcampaignplacementgetEffectiveMessengerPositions) String() string {
	return string(value)
}
