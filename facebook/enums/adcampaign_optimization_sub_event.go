package enums

type AdcampaignOptimizationSubEvent string

const (
	AdcampaignOptimizationSubEventNone                            AdcampaignOptimizationSubEvent = "NONE"
	AdcampaignOptimizationSubEventPostInteraction                 AdcampaignOptimizationSubEvent = "POST_INTERACTION"
	AdcampaignOptimizationSubEventTravelIntent                    AdcampaignOptimizationSubEvent = "TRAVEL_INTENT"
	AdcampaignOptimizationSubEventTravelIntentBucketX01           AdcampaignOptimizationSubEvent = "TRAVEL_INTENT_BUCKET_01"
	AdcampaignOptimizationSubEventTravelIntentBucketX02           AdcampaignOptimizationSubEvent = "TRAVEL_INTENT_BUCKET_02"
	AdcampaignOptimizationSubEventTravelIntentBucketX03           AdcampaignOptimizationSubEvent = "TRAVEL_INTENT_BUCKET_03"
	AdcampaignOptimizationSubEventTravelIntentBucketX04           AdcampaignOptimizationSubEvent = "TRAVEL_INTENT_BUCKET_04"
	AdcampaignOptimizationSubEventTravelIntentBucketX05           AdcampaignOptimizationSubEvent = "TRAVEL_INTENT_BUCKET_05"
	AdcampaignOptimizationSubEventTravelIntentNoDestinationIntent AdcampaignOptimizationSubEvent = "TRAVEL_INTENT_NO_DESTINATION_INTENT"
	AdcampaignOptimizationSubEventTripConsideration               AdcampaignOptimizationSubEvent = "TRIP_CONSIDERATION"
	AdcampaignOptimizationSubEventVideoSoundOn                    AdcampaignOptimizationSubEvent = "VIDEO_SOUND_ON"
)

func (value AdcampaignOptimizationSubEvent) String() string {
	return string(value)
}
