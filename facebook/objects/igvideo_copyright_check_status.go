package objects

type IGVideoCopyrightCheckStatus struct {
	MatchesFound *bool   `json:"matches_found,omitempty"`
	Status       *string `json:"status,omitempty"`
}
