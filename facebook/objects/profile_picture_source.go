package objects

type ProfilePictureSource struct {
	Bottom       *uint64 `json:"bottom,omitempty"`
	CacheKey     *string `json:"cache_key,omitempty"`
	Height       *uint64 `json:"height,omitempty"`
	IsSilhouette *bool   `json:"is_silhouette,omitempty"`
	Left         *uint64 `json:"left,omitempty"`
	Right        *uint64 `json:"right,omitempty"`
	Top          *uint64 `json:"top,omitempty"`
	URL          *string `json:"url,omitempty"`
	Width        *uint64 `json:"width,omitempty"`
}

var ProfilePictureSourceFields = struct {
	Bottom       string
	CacheKey     string
	Height       string
	IsSilhouette string
	Left         string
	Right        string
	Top          string
	URL          string
	Width        string
}{
	Bottom:       "bottom",
	CacheKey:     "cache_key",
	Height:       "height",
	IsSilhouette: "is_silhouette",
	Left:         "left",
	Right:        "right",
	Top:          "top",
	URL:          "url",
	Width:        "width",
}
