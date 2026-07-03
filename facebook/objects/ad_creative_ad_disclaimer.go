package objects

type AdCreativeAdDisclaimer struct {
	Text  *string `json:"text,omitempty"`
	Title *string `json:"title,omitempty"`
	URL   *string `json:"url,omitempty"`
}

var AdCreativeAdDisclaimerFields = struct {
	Text  string
	Title string
	URL   string
}{
	Text:  "text",
	Title: "title",
	URL:   "url",
}
