package objects

type AgeRange struct {
	Max *uint64 `json:"max,omitempty"`
	Min *uint64 `json:"min,omitempty"`
}
