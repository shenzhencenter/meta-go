package objects

type AdKeywords struct {
	Brands            *[]string `json:"brands,omitempty"`
	ProductCategories *[]string `json:"product_categories,omitempty"`
	ProductNames      *[]string `json:"product_names,omitempty"`
	SearchTerms       *[]string `json:"search_terms,omitempty"`
}

var AdKeywordsFields = struct {
	Brands            string
	ProductCategories string
	ProductNames      string
	SearchTerms       string
}{
	Brands:            "brands",
	ProductCategories: "product_categories",
	ProductNames:      "product_names",
	SearchTerms:       "search_terms",
}
