package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type WebPublisher struct {
	DomainURL     *string  `json:"domain_url,omitempty"`
	ID            *core.ID `json:"id,omitempty"`
	PublisherName *string  `json:"publisher_name,omitempty"`
}
