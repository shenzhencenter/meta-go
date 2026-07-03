package objects

type CommerceSettings struct {
	Inventory      *int `json:"inventory,omitempty"`
	TotalInventory *int `json:"total_inventory,omitempty"`
}
