package objects

type FundingSourceDetailsCouponTiering struct {
	CouponTieringNew          *map[string]interface{} `json:"coupon_tiering_new,omitempty"`
	CouponTieringReactivation *map[string]interface{} `json:"coupon_tiering_reactivation,omitempty"`
}
