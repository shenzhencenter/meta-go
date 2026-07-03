package enums

type AdcampaigninsightsActionBreakdownsEnumParam string

const (
	AdcampaigninsightsActionBreakdownsEnumParamActionCanvasComponentName AdcampaigninsightsActionBreakdownsEnumParam = "action_canvas_component_name"
	AdcampaigninsightsActionBreakdownsEnumParamActionCarouselCardID      AdcampaigninsightsActionBreakdownsEnumParam = "action_carousel_card_id"
	AdcampaigninsightsActionBreakdownsEnumParamActionCarouselCardName    AdcampaigninsightsActionBreakdownsEnumParam = "action_carousel_card_name"
	AdcampaigninsightsActionBreakdownsEnumParamActionDestination         AdcampaigninsightsActionBreakdownsEnumParam = "action_destination"
	AdcampaigninsightsActionBreakdownsEnumParamActionDevice              AdcampaigninsightsActionBreakdownsEnumParam = "action_device"
	AdcampaigninsightsActionBreakdownsEnumParamActionReaction            AdcampaigninsightsActionBreakdownsEnumParam = "action_reaction"
	AdcampaigninsightsActionBreakdownsEnumParamActionTargetID            AdcampaigninsightsActionBreakdownsEnumParam = "action_target_id"
	AdcampaigninsightsActionBreakdownsEnumParamActionType                AdcampaigninsightsActionBreakdownsEnumParam = "action_type"
	AdcampaigninsightsActionBreakdownsEnumParamActionVideoSound          AdcampaigninsightsActionBreakdownsEnumParam = "action_video_sound"
	AdcampaigninsightsActionBreakdownsEnumParamActionVideoType           AdcampaigninsightsActionBreakdownsEnumParam = "action_video_type"
	AdcampaigninsightsActionBreakdownsEnumParamConversionDestination     AdcampaigninsightsActionBreakdownsEnumParam = "conversion_destination"
	AdcampaigninsightsActionBreakdownsEnumParamIsBusinessAiAssisted      AdcampaigninsightsActionBreakdownsEnumParam = "is_business_ai_assisted"
	AdcampaigninsightsActionBreakdownsEnumParamMatchedPersonaID          AdcampaigninsightsActionBreakdownsEnumParam = "matched_persona_id"
	AdcampaigninsightsActionBreakdownsEnumParamMatchedPersonaName        AdcampaigninsightsActionBreakdownsEnumParam = "matched_persona_name"
	AdcampaigninsightsActionBreakdownsEnumParamSignalSourceBucket        AdcampaigninsightsActionBreakdownsEnumParam = "signal_source_bucket"
	AdcampaigninsightsActionBreakdownsEnumParamStandardEventContentType  AdcampaigninsightsActionBreakdownsEnumParam = "standard_event_content_type"
)

func (value AdcampaigninsightsActionBreakdownsEnumParam) String() string {
	return string(value)
}
