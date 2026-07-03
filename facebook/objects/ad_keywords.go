package objects

type AdKeywords struct {
	Brands            *[]string `json:"brands,omitempty"`
	ProductCategories *[]string `json:"product_categories,omitempty"`
	ProductNames      *[]string `json:"product_names,omitempty"`
	SearchTerms       *[]string `json:"search_terms,omitempty"`
}
