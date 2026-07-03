package objects

type AdDynamicCreative struct {
	PreviewURL *string `json:"preview_url,omitempty"`
}

var AdDynamicCreativeFields = struct {
	PreviewURL string
}{
	PreviewURL: "preview_url",
}
