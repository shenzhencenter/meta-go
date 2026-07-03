package objects

type PlatformImageSource struct {
	Height *uint64 `json:"height,omitempty"`
	Source *string `json:"source,omitempty"`
	Width  *uint64 `json:"width,omitempty"`
}

var PlatformImageSourceFields = struct {
	Height string
	Source string
	Width  string
}{
	Height: "height",
	Source: "source",
	Width:  "width",
}
