package objects

type UserLeadGenFieldData struct {
	Name   *string   `json:"name,omitempty"`
	Values *[]string `json:"values,omitempty"`
}

var UserLeadGenFieldDataFields = struct {
	Name   string
	Values string
}{
	Name:   "name",
	Values: "values",
}
