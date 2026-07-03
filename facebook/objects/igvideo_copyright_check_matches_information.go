package objects

type IGVideoCopyrightCheckMatchesInformation struct {
	CopyrightMatches *[]map[string]interface{}    `json:"copyright_matches,omitempty"`
	Status           *IGVideoCopyrightCheckStatus `json:"status,omitempty"`
}

var IGVideoCopyrightCheckMatchesInformationFields = struct {
	CopyrightMatches string
	Status           string
}{
	CopyrightMatches: "copyright_matches",
	Status:           "status",
}
