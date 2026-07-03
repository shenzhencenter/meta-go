package objects

type UserDevice struct {
	Hardware *string `json:"hardware,omitempty"`
	Os       *string `json:"os,omitempty"`
}

var UserDeviceFields = struct {
	Hardware string
	Os       string
}{
	Hardware: "hardware",
	Os:       "os",
}
