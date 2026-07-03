package objects

type LeadGenClientValidationRules struct {
	ExcludeEmojiAndSpecialCharsEnabled *bool `json:"exclude_emoji_and_special_chars_enabled,omitempty"`
	MaxLengthValue                     *int  `json:"max_length_value,omitempty"`
	MinLengthValue                     *int  `json:"min_length_value,omitempty"`
}
