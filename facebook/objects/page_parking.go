package objects

type PageParking struct {
	Lot    *uint64 `json:"lot,omitempty"`
	Street *uint64 `json:"street,omitempty"`
	Valet  *uint64 `json:"valet,omitempty"`
}
