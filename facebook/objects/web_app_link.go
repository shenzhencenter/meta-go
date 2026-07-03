package objects

type WebAppLink struct {
	ShouldFallback *bool   `json:"should_fallback,omitempty"`
	URL            *string `json:"url,omitempty"`
}

var WebAppLinkFields = struct {
	ShouldFallback string
	URL            string
}{
	ShouldFallback: "should_fallback",
	URL:            "url",
}
