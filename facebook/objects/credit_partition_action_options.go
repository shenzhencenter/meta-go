package objects

type CreditPartitionActionOptions struct {
	LiabilityType *map[string]interface{} `json:"liability_type,omitempty"`
	PartitionType *map[string]interface{} `json:"partition_type,omitempty"`
	SendBillTo    *map[string]interface{} `json:"send_bill_to,omitempty"`
}

var CreditPartitionActionOptionsFields = struct {
	LiabilityType string
	PartitionType string
	SendBillTo    string
}{
	LiabilityType: "liability_type",
	PartitionType: "partition_type",
	SendBillTo:    "send_bill_to",
}
