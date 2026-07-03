package objects

type VideoCopyrightCheckStatus struct {
	MatchesFound *bool   `json:"matches_found,omitempty"`
	Status       *string `json:"status,omitempty"`
}
