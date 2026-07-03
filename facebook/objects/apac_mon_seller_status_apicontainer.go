package objects

type ApacMonSellerStatusAPIContainer struct {
	StructuredMessagingCommerce *map[string]interface{} `json:"structured_messaging_commerce,omitempty"`
}

var ApacMonSellerStatusAPIContainerFields = struct {
	StructuredMessagingCommerce string
}{
	StructuredMessagingCommerce: "structured_messaging_commerce",
}
