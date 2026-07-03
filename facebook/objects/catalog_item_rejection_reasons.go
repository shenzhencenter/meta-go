package objects

type CatalogItemRejectionReasons struct {
	Capability           *string                   `json:"capability,omitempty"`
	RejectionInformation *[]map[string]interface{} `json:"rejection_information,omitempty"`
}

var CatalogItemRejectionReasonsFields = struct {
	Capability           string
	RejectionInformation string
}{
	Capability:           "capability",
	RejectionInformation: "rejection_information",
}
