package objects

type VideoCopyrightCheckStatus struct {
	MatchesFound *bool   `json:"matches_found,omitempty"`
	Status       *string `json:"status,omitempty"`
}

var VideoCopyrightCheckStatusFields = struct {
	MatchesFound string
	Status       string
}{
	MatchesFound: "matches_found",
	Status:       "status",
}
