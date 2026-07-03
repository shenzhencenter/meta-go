package objects

type AgeRange struct {
	Max *uint64 `json:"max,omitempty"`
	Min *uint64 `json:"min,omitempty"`
}

var AgeRangeFields = struct {
	Max string
	Min string
}{
	Max: "max",
	Min: "min",
}
