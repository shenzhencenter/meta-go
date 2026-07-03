package objects

type UserDevice struct {
	Hardware *string `json:"hardware,omitempty"`
	Os       *string `json:"os,omitempty"`
}
