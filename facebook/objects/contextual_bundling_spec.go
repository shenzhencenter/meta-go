package objects

type ContextualBundlingSpec struct {
	Status *string `json:"status,omitempty"`
}

var ContextualBundlingSpecFields = struct {
	Status string
}{
	Status: "status",
}
