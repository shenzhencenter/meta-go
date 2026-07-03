package enums

type AdaccountadsetsOptimizationSubEventEnumParam string

const (
	AdaccountadsetsOptimizationSubEventEnumParamNone                            AdaccountadsetsOptimizationSubEventEnumParam = "NONE"
	AdaccountadsetsOptimizationSubEventEnumParamPostInteraction                 AdaccountadsetsOptimizationSubEventEnumParam = "POST_INTERACTION"
	AdaccountadsetsOptimizationSubEventEnumParamTravelIntent                    AdaccountadsetsOptimizationSubEventEnumParam = "TRAVEL_INTENT"
	AdaccountadsetsOptimizationSubEventEnumParamTravelIntentBucketX01           AdaccountadsetsOptimizationSubEventEnumParam = "TRAVEL_INTENT_BUCKET_01"
	AdaccountadsetsOptimizationSubEventEnumParamTravelIntentBucketX02           AdaccountadsetsOptimizationSubEventEnumParam = "TRAVEL_INTENT_BUCKET_02"
	AdaccountadsetsOptimizationSubEventEnumParamTravelIntentBucketX03           AdaccountadsetsOptimizationSubEventEnumParam = "TRAVEL_INTENT_BUCKET_03"
	AdaccountadsetsOptimizationSubEventEnumParamTravelIntentBucketX04           AdaccountadsetsOptimizationSubEventEnumParam = "TRAVEL_INTENT_BUCKET_04"
	AdaccountadsetsOptimizationSubEventEnumParamTravelIntentBucketX05           AdaccountadsetsOptimizationSubEventEnumParam = "TRAVEL_INTENT_BUCKET_05"
	AdaccountadsetsOptimizationSubEventEnumParamTravelIntentNoDestinationIntent AdaccountadsetsOptimizationSubEventEnumParam = "TRAVEL_INTENT_NO_DESTINATION_INTENT"
	AdaccountadsetsOptimizationSubEventEnumParamTripConsideration               AdaccountadsetsOptimizationSubEventEnumParam = "TRIP_CONSIDERATION"
	AdaccountadsetsOptimizationSubEventEnumParamVideoSoundOn                    AdaccountadsetsOptimizationSubEventEnumParam = "VIDEO_SOUND_ON"
)

func (value AdaccountadsetsOptimizationSubEventEnumParam) String() string {
	return string(value)
}
