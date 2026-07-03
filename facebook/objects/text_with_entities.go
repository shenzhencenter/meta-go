package objects

type TextWithEntities struct {
	Text *string `json:"text,omitempty"`
}

var TextWithEntitiesFields = struct {
	Text string
}{
	Text: "text",
}
