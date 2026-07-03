package objects

type AuthLink struct {
	Link *string `json:"link,omitempty"`
}

var AuthLinkFields = struct {
	Link string
}{
	Link: "link",
}
