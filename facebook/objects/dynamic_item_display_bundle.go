package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type DynamicItemDisplayBundle struct {
	AdditionalUrls *[]map[string]string `json:"additional_urls,omitempty"`
	Description    *string              `json:"description,omitempty"`
	ID             *core.ID             `json:"id,omitempty"`
	Name           *string              `json:"name,omitempty"`
	ProductSet     *ProductSet          `json:"product_set,omitempty"`
	TextTokens     *[]map[string]string `json:"text_tokens,omitempty"`
	URL            *string              `json:"url,omitempty"`
}
