package objects

type PageThreadOwner struct {
	ThreadOwner *map[string]interface{} `json:"thread_owner,omitempty"`
}

var PageThreadOwnerFields = struct {
	ThreadOwner string
}{
	ThreadOwner: "thread_owner",
}
