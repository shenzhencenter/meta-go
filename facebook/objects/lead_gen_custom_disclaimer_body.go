package objects

type LeadGenCustomDisclaimerBody struct {
	Text        *string                     `json:"text,omitempty"`
	URLEntities *[]LeadGenURLEntityAtRanges `json:"url_entities,omitempty"`
}
