package objects

type PlatformImageSource struct {
	Height *uint64 `json:"height,omitempty"`
	Source *string `json:"source,omitempty"`
	Width  *uint64 `json:"width,omitempty"`
}
