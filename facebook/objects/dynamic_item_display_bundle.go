package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var DynamicItemDisplayBundleFields = struct {
	AdditionalUrls string
	Description    string
	ID             string
	Name           string
	ProductSet     string
	TextTokens     string
	URL            string
}{
	AdditionalUrls: "additional_urls",
	Description:    "description",
	ID:             "id",
	Name:           "name",
	ProductSet:     "product_set",
	TextTokens:     "text_tokens",
	URL:            "url",
}
