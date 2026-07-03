package objects

type PageRestaurantSpecialties struct {
	Breakfast *uint64 `json:"breakfast,omitempty"`
	Coffee    *uint64 `json:"coffee,omitempty"`
	Dinner    *uint64 `json:"dinner,omitempty"`
	Drinks    *uint64 `json:"drinks,omitempty"`
	Lunch     *uint64 `json:"lunch,omitempty"`
}

var PageRestaurantSpecialtiesFields = struct {
	Breakfast string
	Coffee    string
	Dinner    string
	Drinks    string
	Lunch     string
}{
	Breakfast: "breakfast",
	Coffee:    "coffee",
	Dinner:    "dinner",
	Drinks:    "drinks",
	Lunch:     "lunch",
}
