package objects

type CanvasPreview struct {
	Body *string `json:"body,omitempty"`
}

var CanvasPreviewFields = struct {
	Body string
}{
	Body: "body",
}
