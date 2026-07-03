package objects

type PublisherDeliveryReport struct {
	ContentTypes         *[]string `json:"content_types,omitempty"`
	EstimatedImpressions *uint64   `json:"estimated_impressions,omitempty"`
	Name                 *string   `json:"name,omitempty"`
	Status               *string   `json:"status,omitempty"`
	URL                  *string   `json:"url,omitempty"`
}

var PublisherDeliveryReportFields = struct {
	ContentTypes         string
	EstimatedImpressions string
	Name                 string
	Status               string
	URL                  string
}{
	ContentTypes:         "content_types",
	EstimatedImpressions: "estimated_impressions",
	Name:                 "name",
	Status:               "status",
	URL:                  "url",
}
