package objects

type AdsEligibility struct {
	LiveShopping *map[string]interface{} `json:"live_shopping,omitempty"`
}

var AdsEligibilityFields = struct {
	LiveShopping string
}{
	LiveShopping: "live_shopping",
}
