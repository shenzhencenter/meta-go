package objects

type PaymentPricepoints struct {
	Mobile *[]map[string]interface{} `json:"mobile,omitempty"`
}

var PaymentPricepointsFields = struct {
	Mobile string
}{
	Mobile: "mobile",
}
