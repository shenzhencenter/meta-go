package objects

type VideoStatusProcessingPhase struct {
	Errors *[]VideoStatusError `json:"errors,omitempty"`
	Status *string             `json:"status,omitempty"`
}
