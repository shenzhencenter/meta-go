package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdvertiserVerificationStatus struct {
	BannerType         *string    `json:"banner_type,omitempty"`
	GracePeriodEndsAt  *core.Time `json:"grace_period_ends_at,omitempty"`
	UfacRedirectURI    *string    `json:"ufac_redirect_uri,omitempty"`
	VerificationStatus *string    `json:"verification_status,omitempty"`
}

var AdvertiserVerificationStatusFields = struct {
	BannerType         string
	GracePeriodEndsAt  string
	UfacRedirectURI    string
	VerificationStatus string
}{
	BannerType:         "banner_type",
	GracePeriodEndsAt:  "grace_period_ends_at",
	UfacRedirectURI:    "ufac_redirect_uri",
	VerificationStatus: "verification_status",
}
