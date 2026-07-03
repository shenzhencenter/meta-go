package objects

type Engagement struct {
	Count                     *uint64 `json:"count,omitempty"`
	CountString               *string `json:"count_string,omitempty"`
	CountStringWithLike       *string `json:"count_string_with_like,omitempty"`
	CountStringWithoutLike    *string `json:"count_string_without_like,omitempty"`
	SocialSentence            *string `json:"social_sentence,omitempty"`
	SocialSentenceWithLike    *string `json:"social_sentence_with_like,omitempty"`
	SocialSentenceWithoutLike *string `json:"social_sentence_without_like,omitempty"`
}
