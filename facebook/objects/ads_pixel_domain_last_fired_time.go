package objects

type AdsPixelDomainLastFiredTime struct {
	DomainName    *string `json:"domain_name,omitempty"`
	LastFiredTime *int    `json:"last_fired_time,omitempty"`
}

var AdsPixelDomainLastFiredTimeFields = struct {
	DomainName    string
	LastFiredTime string
}{
	DomainName:    "domain_name",
	LastFiredTime: "last_fired_time",
}
