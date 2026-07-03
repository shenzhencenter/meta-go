package objects

type CanvasBodyElement struct {
	Element *map[string]interface{} `json:"element,omitempty"`
}

var CanvasBodyElementFields = struct {
	Element string
}{
	Element: "element",
}
