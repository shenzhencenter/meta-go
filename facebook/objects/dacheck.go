package objects

type DACheck struct {
	ActionURI   *string `json:"action_uri,omitempty"`
	Description *string `json:"description,omitempty"`
	Key         *string `json:"key,omitempty"`
	Result      *string `json:"result,omitempty"`
	Title       *string `json:"title,omitempty"`
	UserMessage *string `json:"user_message,omitempty"`
}

var DACheckFields = struct {
	ActionURI   string
	Description string
	Key         string
	Result      string
	Title       string
	UserMessage string
}{
	ActionURI:   "action_uri",
	Description: "description",
	Key:         "key",
	Result:      "result",
	Title:       "title",
	UserMessage: "user_message",
}
