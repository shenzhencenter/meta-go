package objects

type Privacy struct {
	Allow       *string `json:"allow,omitempty"`
	Deny        *string `json:"deny,omitempty"`
	Description *string `json:"description,omitempty"`
	Friends     *string `json:"friends,omitempty"`
	Networks    *string `json:"networks,omitempty"`
	Value       *string `json:"value,omitempty"`
}
