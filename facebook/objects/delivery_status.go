package objects

type DeliveryStatus struct {
	Status      *string   `json:"status,omitempty"`
	Substatuses *[]string `json:"substatuses,omitempty"`
}
