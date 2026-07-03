package enums

type AdaccountinsightsActionBreakdownsEnumParam string

const (
	AdaccountinsightsActionBreakdownsEnumParamActionCanvasComponentName AdaccountinsightsActionBreakdownsEnumParam = "action_canvas_component_name"
	AdaccountinsightsActionBreakdownsEnumParamActionCarouselCardID      AdaccountinsightsActionBreakdownsEnumParam = "action_carousel_card_id"
	AdaccountinsightsActionBreakdownsEnumParamActionCarouselCardName    AdaccountinsightsActionBreakdownsEnumParam = "action_carousel_card_name"
	AdaccountinsightsActionBreakdownsEnumParamActionDestination         AdaccountinsightsActionBreakdownsEnumParam = "action_destination"
	AdaccountinsightsActionBreakdownsEnumParamActionDevice              AdaccountinsightsActionBreakdownsEnumParam = "action_device"
	AdaccountinsightsActionBreakdownsEnumParamActionReaction            AdaccountinsightsActionBreakdownsEnumParam = "action_reaction"
	AdaccountinsightsActionBreakdownsEnumParamActionTargetID            AdaccountinsightsActionBreakdownsEnumParam = "action_target_id"
	AdaccountinsightsActionBreakdownsEnumParamActionType                AdaccountinsightsActionBreakdownsEnumParam = "action_type"
	AdaccountinsightsActionBreakdownsEnumParamActionVideoSound          AdaccountinsightsActionBreakdownsEnumParam = "action_video_sound"
	AdaccountinsightsActionBreakdownsEnumParamActionVideoType           AdaccountinsightsActionBreakdownsEnumParam = "action_video_type"
	AdaccountinsightsActionBreakdownsEnumParamConversionDestination     AdaccountinsightsActionBreakdownsEnumParam = "conversion_destination"
	AdaccountinsightsActionBreakdownsEnumParamIsBusinessAiAssisted      AdaccountinsightsActionBreakdownsEnumParam = "is_business_ai_assisted"
	AdaccountinsightsActionBreakdownsEnumParamMatchedPersonaID          AdaccountinsightsActionBreakdownsEnumParam = "matched_persona_id"
	AdaccountinsightsActionBreakdownsEnumParamMatchedPersonaName        AdaccountinsightsActionBreakdownsEnumParam = "matched_persona_name"
	AdaccountinsightsActionBreakdownsEnumParamSignalSourceBucket        AdaccountinsightsActionBreakdownsEnumParam = "signal_source_bucket"
	AdaccountinsightsActionBreakdownsEnumParamStandardEventContentType  AdaccountinsightsActionBreakdownsEnumParam = "standard_event_content_type"
)

func (value AdaccountinsightsActionBreakdownsEnumParam) String() string {
	return string(value)
}
