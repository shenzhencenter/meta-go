package objects

type AdEntityTargetSpend struct {
	Amount     *string `json:"amount,omitempty"`
	HasError   *bool   `json:"has_error,omitempty"`
	IsAccurate *bool   `json:"is_accurate,omitempty"`
	IsProrated *bool   `json:"is_prorated,omitempty"`
	IsUpdating *bool   `json:"is_updating,omitempty"`
}

var AdEntityTargetSpendFields = struct {
	Amount     string
	HasError   string
	IsAccurate string
	IsProrated string
	IsUpdating string
}{
	Amount:     "amount",
	HasError:   "has_error",
	IsAccurate: "is_accurate",
	IsProrated: "is_prorated",
	IsUpdating: "is_updating",
}
