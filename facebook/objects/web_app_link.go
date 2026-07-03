package objects

type WebAppLink struct {
	ShouldFallback *bool   `json:"should_fallback,omitempty"`
	URL            *string `json:"url,omitempty"`
}
