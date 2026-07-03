package objects

type CommerceSettings struct {
	Inventory      *int `json:"inventory,omitempty"`
	TotalInventory *int `json:"total_inventory,omitempty"`
}

var CommerceSettingsFields = struct {
	Inventory      string
	TotalInventory string
}{
	Inventory:      "inventory",
	TotalInventory: "total_inventory",
}
