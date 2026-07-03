package objects

type FinanceObject struct {
	FinancePermission *string                 `json:"finance_permission,omitempty"`
	User              *map[string]interface{} `json:"user,omitempty"`
}

var FinanceObjectFields = struct {
	FinancePermission string
	User              string
}{
	FinancePermission: "finance_permission",
	User:              "user",
}
