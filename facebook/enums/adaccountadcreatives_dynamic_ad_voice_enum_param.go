package enums

type AdaccountadcreativesDynamicAdVoiceEnumParam string

const (
	AdaccountadcreativesDynamicAdVoiceEnumParamDynamic    AdaccountadcreativesDynamicAdVoiceEnumParam = "DYNAMIC"
	AdaccountadcreativesDynamicAdVoiceEnumParamStoryOwner AdaccountadcreativesDynamicAdVoiceEnumParam = "STORY_OWNER"
)

func (value AdaccountadcreativesDynamicAdVoiceEnumParam) String() string {
	return string(value)
}
