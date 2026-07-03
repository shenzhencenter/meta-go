package objects

type PageCTXDefaultGreetingText struct {
	Ctd  *string `json:"ctd,omitempty"`
	Ctm  *string `json:"ctm,omitempty"`
	Ctwa *string `json:"ctwa,omitempty"`
}

var PageCTXDefaultGreetingTextFields = struct {
	Ctd  string
	Ctm  string
	Ctwa string
}{
	Ctd:  "ctd",
	Ctm:  "ctm",
	Ctwa: "ctwa",
}
