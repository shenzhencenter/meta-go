package objects

type AdCreativeDestinationSpec struct {
	DestinationType          *string                 `json:"destination_type,omitempty"`
	NativeCommerceExperience *map[string]interface{} `json:"native_commerce_experience,omitempty"`
	Website                  *map[string]interface{} `json:"website,omitempty"`
}

var AdCreativeDestinationSpecFields = struct {
	DestinationType          string
	NativeCommerceExperience string
	Website                  string
}{
	DestinationType:          "destination_type",
	NativeCommerceExperience: "native_commerce_experience",
	Website:                  "website",
}
