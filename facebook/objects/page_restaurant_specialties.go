package objects

type PageRestaurantSpecialties struct {
	Breakfast *uint64 `json:"breakfast,omitempty"`
	Coffee    *uint64 `json:"coffee,omitempty"`
	Dinner    *uint64 `json:"dinner,omitempty"`
	Drinks    *uint64 `json:"drinks,omitempty"`
	Lunch     *uint64 `json:"lunch,omitempty"`
}
