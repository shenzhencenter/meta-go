package objects

type InstagramShoppingMerchantReviewMessage struct {
	HelpURL *string `json:"help_url,omitempty"`
	Message *string `json:"message,omitempty"`
}

var InstagramShoppingMerchantReviewMessageFields = struct {
	HelpURL string
	Message string
}{
	HelpURL: "help_url",
	Message: "message",
}
