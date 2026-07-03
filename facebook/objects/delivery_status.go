package objects

type DeliveryStatus struct {
	Status      *string   `json:"status,omitempty"`
	Substatuses *[]string `json:"substatuses,omitempty"`
}

var DeliveryStatusFields = struct {
	Status      string
	Substatuses string
}{
	Status:      "status",
	Substatuses: "substatuses",
}
