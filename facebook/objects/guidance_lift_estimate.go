package objects

type GuidanceLiftEstimate struct {
	ActualX7dCpr    *float64 `json:"actual_7d_cpr,omitempty"`
	AdoptionDate    *string  `json:"adoption_date,omitempty"`
	GuidanceName    *string  `json:"guidance_name,omitempty"`
	LiftEstimation  *float64 `json:"lift_estimation,omitempty"`
	PredictedX7dCpr *float64 `json:"predicted_7d_cpr,omitempty"`
}

var GuidanceLiftEstimateFields = struct {
	ActualX7dCpr    string
	AdoptionDate    string
	GuidanceName    string
	LiftEstimation  string
	PredictedX7dCpr string
}{
	ActualX7dCpr:    "actual_7d_cpr",
	AdoptionDate:    "adoption_date",
	GuidanceName:    "guidance_name",
	LiftEstimation:  "lift_estimation",
	PredictedX7dCpr: "predicted_7d_cpr",
}
