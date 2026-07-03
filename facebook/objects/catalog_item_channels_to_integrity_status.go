package objects

type CatalogItemChannelsToIntegrityStatus struct {
	Channels             *[]string               `json:"channels,omitempty"`
	RejectionInformation *map[string]interface{} `json:"rejection_information,omitempty"`
}

var CatalogItemChannelsToIntegrityStatusFields = struct {
	Channels             string
	RejectionInformation string
}{
	Channels:             "channels",
	RejectionInformation: "rejection_information",
}
