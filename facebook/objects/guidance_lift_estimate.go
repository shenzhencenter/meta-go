package objects

type GuidanceLiftEstimate struct {
	ActualX7dCpr    *float64 `json:"actual_7d_cpr,omitempty"`
	AdoptionDate    *string  `json:"adoption_date,omitempty"`
	GuidanceName    *string  `json:"guidance_name,omitempty"`
	LiftEstimation  *float64 `json:"lift_estimation,omitempty"`
	PredictedX7dCpr *float64 `json:"predicted_7d_cpr,omitempty"`
}
