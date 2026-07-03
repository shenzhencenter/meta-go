package objects

type AdCampaignConversionValueExpressionSpec struct {
	AdjustmentSign   *string `json:"adjustment_sign,omitempty"`
	AdjustmentWeight *uint64 `json:"adjustment_weight,omitempty"`
	DestinationType  *string `json:"destination_type,omitempty"`
}
