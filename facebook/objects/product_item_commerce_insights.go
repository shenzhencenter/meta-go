package objects

type ProductItemCommerceInsights struct {
	MessageSends       *uint64 `json:"message_sends,omitempty"`
	OrganicImpressions *uint64 `json:"organic_impressions,omitempty"`
	PaidImpressions    *uint64 `json:"paid_impressions,omitempty"`
}
