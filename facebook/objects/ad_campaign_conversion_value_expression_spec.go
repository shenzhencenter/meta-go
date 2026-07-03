package objects

type AdCampaignConversionValueExpressionSpec struct {
	AdjustmentSign   *string `json:"adjustment_sign,omitempty"`
	AdjustmentWeight *uint64 `json:"adjustment_weight,omitempty"`
	DestinationType  *string `json:"destination_type,omitempty"`
}

var AdCampaignConversionValueExpressionSpecFields = struct {
	AdjustmentSign   string
	AdjustmentWeight string
	DestinationType  string
}{
	AdjustmentSign:   "adjustment_sign",
	AdjustmentWeight: "adjustment_weight",
	DestinationType:  "destination_type",
}
