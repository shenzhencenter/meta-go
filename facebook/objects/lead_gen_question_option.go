package objects

type LeadGenQuestionOption struct {
	Key   *string `json:"key,omitempty"`
	Photo *Photo  `json:"photo,omitempty"`
	Value *string `json:"value,omitempty"`
}
