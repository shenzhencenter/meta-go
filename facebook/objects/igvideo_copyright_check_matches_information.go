package objects

type IGVideoCopyrightCheckMatchesInformation struct {
	CopyrightMatches *[]map[string]interface{}    `json:"copyright_matches,omitempty"`
	Status           *IGVideoCopyrightCheckStatus `json:"status,omitempty"`
}
