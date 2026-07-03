package objects

type VideoStatusProcessingPhase struct {
	Errors *[]VideoStatusError `json:"errors,omitempty"`
	Status *string             `json:"status,omitempty"`
}

var VideoStatusProcessingPhaseFields = struct {
	Errors string
	Status string
}{
	Errors: "errors",
	Status: "status",
}
