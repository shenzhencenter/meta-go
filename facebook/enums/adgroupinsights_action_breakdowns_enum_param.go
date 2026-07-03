package enums

type AdgroupinsightsActionBreakdownsEnumParam string

const (
	AdgroupinsightsActionBreakdownsEnumParamActionCanvasComponentName AdgroupinsightsActionBreakdownsEnumParam = "action_canvas_component_name"
	AdgroupinsightsActionBreakdownsEnumParamActionCarouselCardID      AdgroupinsightsActionBreakdownsEnumParam = "action_carousel_card_id"
	AdgroupinsightsActionBreakdownsEnumParamActionCarouselCardName    AdgroupinsightsActionBreakdownsEnumParam = "action_carousel_card_name"
	AdgroupinsightsActionBreakdownsEnumParamActionDestination         AdgroupinsightsActionBreakdownsEnumParam = "action_destination"
	AdgroupinsightsActionBreakdownsEnumParamActionDevice              AdgroupinsightsActionBreakdownsEnumParam = "action_device"
	AdgroupinsightsActionBreakdownsEnumParamActionReaction            AdgroupinsightsActionBreakdownsEnumParam = "action_reaction"
	AdgroupinsightsActionBreakdownsEnumParamActionTargetID            AdgroupinsightsActionBreakdownsEnumParam = "action_target_id"
	AdgroupinsightsActionBreakdownsEnumParamActionType                AdgroupinsightsActionBreakdownsEnumParam = "action_type"
	AdgroupinsightsActionBreakdownsEnumParamActionVideoSound          AdgroupinsightsActionBreakdownsEnumParam = "action_video_sound"
	AdgroupinsightsActionBreakdownsEnumParamActionVideoType           AdgroupinsightsActionBreakdownsEnumParam = "action_video_type"
	AdgroupinsightsActionBreakdownsEnumParamConversionDestination     AdgroupinsightsActionBreakdownsEnumParam = "conversion_destination"
	AdgroupinsightsActionBreakdownsEnumParamIsBusinessAiAssisted      AdgroupinsightsActionBreakdownsEnumParam = "is_business_ai_assisted"
	AdgroupinsightsActionBreakdownsEnumParamMatchedPersonaID          AdgroupinsightsActionBreakdownsEnumParam = "matched_persona_id"
	AdgroupinsightsActionBreakdownsEnumParamMatchedPersonaName        AdgroupinsightsActionBreakdownsEnumParam = "matched_persona_name"
	AdgroupinsightsActionBreakdownsEnumParamSignalSourceBucket        AdgroupinsightsActionBreakdownsEnumParam = "signal_source_bucket"
	AdgroupinsightsActionBreakdownsEnumParamStandardEventContentType  AdgroupinsightsActionBreakdownsEnumParam = "standard_event_content_type"
)

func (value AdgroupinsightsActionBreakdownsEnumParam) String() string {
	return string(value)
}
