package objects

type AdAccountBankInfoList struct {
	Banks *[]map[string]interface{} `json:"banks,omitempty"`
}

var AdAccountBankInfoListFields = struct {
	Banks string
}{
	Banks: "banks",
}
