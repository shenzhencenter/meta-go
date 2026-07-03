package objects

type LeadGenCustomDisclaimerBody struct {
	Text        *string                     `json:"text,omitempty"`
	URLEntities *[]LeadGenURLEntityAtRanges `json:"url_entities,omitempty"`
}

var LeadGenCustomDisclaimerBodyFields = struct {
	Text        string
	URLEntities string
}{
	Text:        "text",
	URLEntities: "url_entities",
}
