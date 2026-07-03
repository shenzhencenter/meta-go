package objects

type ReachFrequencyEstimatesCurve struct {
	Budget            *[]int   `json:"budget,omitempty"`
	Conversion        *[]int   `json:"conversion,omitempty"`
	Impression        *[]int   `json:"impression,omitempty"`
	InterpolatedReach *float64 `json:"interpolated_reach,omitempty"`
	NumPoints         *uint64  `json:"num_points,omitempty"`
	RawImpression     *[]int   `json:"raw_impression,omitempty"`
	RawReach          *[]int   `json:"raw_reach,omitempty"`
	Reach             *[]int   `json:"reach,omitempty"`
}
