package objects

type PageRestaurantServices struct {
	Catering *bool `json:"catering,omitempty"`
	Delivery *bool `json:"delivery,omitempty"`
	Groups   *bool `json:"groups,omitempty"`
	Kids     *bool `json:"kids,omitempty"`
	Outdoor  *bool `json:"outdoor,omitempty"`
	Pickup   *bool `json:"pickup,omitempty"`
	Reserve  *bool `json:"reserve,omitempty"`
	Takeout  *bool `json:"takeout,omitempty"`
	Waiter   *bool `json:"waiter,omitempty"`
	Walkins  *bool `json:"walkins,omitempty"`
}

var PageRestaurantServicesFields = struct {
	Catering string
	Delivery string
	Groups   string
	Kids     string
	Outdoor  string
	Pickup   string
	Reserve  string
	Takeout  string
	Waiter   string
	Walkins  string
}{
	Catering: "catering",
	Delivery: "delivery",
	Groups:   "groups",
	Kids:     "kids",
	Outdoor:  "outdoor",
	Pickup:   "pickup",
	Reserve:  "reserve",
	Takeout:  "takeout",
	Waiter:   "waiter",
	Walkins:  "walkins",
}
