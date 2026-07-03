package objects

type MessengerCallPermissions struct {
	Actions    *[]map[string]interface{} `json:"actions,omitempty"`
	Permission *map[string]interface{}   `json:"permission,omitempty"`
}

var MessengerCallPermissionsFields = struct {
	Actions    string
	Permission string
}{
	Actions:    "actions",
	Permission: "permission",
}
