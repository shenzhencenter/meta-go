package objects

type DACheck struct {
	ActionURI   *string `json:"action_uri,omitempty"`
	Description *string `json:"description,omitempty"`
	Key         *string `json:"key,omitempty"`
	Result      *string `json:"result,omitempty"`
	Title       *string `json:"title,omitempty"`
	UserMessage *string `json:"user_message,omitempty"`
}
