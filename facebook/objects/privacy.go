package objects

type Privacy struct {
	Allow       *string `json:"allow,omitempty"`
	Deny        *string `json:"deny,omitempty"`
	Description *string `json:"description,omitempty"`
	Friends     *string `json:"friends,omitempty"`
	Networks    *string `json:"networks,omitempty"`
	Value       *string `json:"value,omitempty"`
}

var PrivacyFields = struct {
	Allow       string
	Deny        string
	Description string
	Friends     string
	Networks    string
	Value       string
}{
	Allow:       "allow",
	Deny:        "deny",
	Description: "description",
	Friends:     "friends",
	Networks:    "networks",
	Value:       "value",
}
