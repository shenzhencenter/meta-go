package objects

type IGVideoCopyrightCheckStatus struct {
	MatchesFound *bool   `json:"matches_found,omitempty"`
	Status       *string `json:"status,omitempty"`
}

var IGVideoCopyrightCheckStatusFields = struct {
	MatchesFound string
	Status       string
}{
	MatchesFound: "matches_found",
	Status:       "status",
}
