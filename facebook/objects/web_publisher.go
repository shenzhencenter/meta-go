package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WebPublisher struct {
	DomainURL     *string  `json:"domain_url,omitempty"`
	ID            *core.ID `json:"id,omitempty"`
	PublisherName *string  `json:"publisher_name,omitempty"`
}

var WebPublisherFields = struct {
	DomainURL     string
	ID            string
	PublisherName string
}{
	DomainURL:     "domain_url",
	ID:            "id",
	PublisherName: "publisher_name",
}
