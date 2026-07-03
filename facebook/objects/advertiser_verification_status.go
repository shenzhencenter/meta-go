package objects

import (
	"time"
)

type AdvertiserVerificationStatus struct {
	BannerType         *string    `json:"banner_type,omitempty"`
	GracePeriodEndsAt  *time.Time `json:"grace_period_ends_at,omitempty"`
	UfacRedirectURI    *string    `json:"ufac_redirect_uri,omitempty"`
	VerificationStatus *string    `json:"verification_status,omitempty"`
}
