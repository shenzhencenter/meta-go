package objects

type FinanceObject struct {
	FinancePermission *string                 `json:"finance_permission,omitempty"`
	User              *map[string]interface{} `json:"user,omitempty"`
}
