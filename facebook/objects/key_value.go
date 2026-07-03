package objects

type KeyValue struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

var KeyValueFields = struct {
	Key   string
	Value string
}{
	Key:   "key",
	Value: "value",
}
