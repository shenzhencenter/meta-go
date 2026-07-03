package objects

type ProductItemVideoData struct {
	Tags *[]string `json:"tags,omitempty"`
	URL  *string   `json:"url,omitempty"`
}

var ProductItemVideoDataFields = struct {
	Tags string
	URL  string
}{
	Tags: "tags",
	URL:  "url",
}
