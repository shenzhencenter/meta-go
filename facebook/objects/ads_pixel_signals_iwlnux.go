package objects

type AdsPixelSignalsIWLNux struct {
	BackgroundColor *string `json:"background_color,omitempty"`
	Content         *string `json:"content,omitempty"`
	ContentColor    *string `json:"content_color,omitempty"`
	ContentSize     *string `json:"content_size,omitempty"`
	ImgURL          *string `json:"img_url,omitempty"`
}

var AdsPixelSignalsIWLNuxFields = struct {
	BackgroundColor string
	Content         string
	ContentColor    string
	ContentSize     string
	ImgURL          string
}{
	BackgroundColor: "background_color",
	Content:         "content",
	ContentColor:    "content_color",
	ContentSize:     "content_size",
	ImgURL:          "img_url",
}
