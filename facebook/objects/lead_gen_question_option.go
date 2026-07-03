package objects

type LeadGenQuestionOption struct {
	Key   *string `json:"key,omitempty"`
	Photo *Photo  `json:"photo,omitempty"`
	Value *string `json:"value,omitempty"`
}

var LeadGenQuestionOptionFields = struct {
	Key   string
	Photo string
	Value string
}{
	Key:   "key",
	Photo: "photo",
	Value: "value",
}
