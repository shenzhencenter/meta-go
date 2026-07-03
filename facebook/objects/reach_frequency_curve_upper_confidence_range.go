package objects

type ReachFrequencyCurveUpperConfidenceRange struct {
	ImpressionUpper        *[]int  `json:"impression_upper,omitempty"`
	NumPoints              *uint64 `json:"num_points,omitempty"`
	Reach                  *[]int  `json:"reach,omitempty"`
	ReachUpper             *[]int  `json:"reach_upper,omitempty"`
	UniqVideoViewsX2sUpper *[]int  `json:"uniq_video_views_2s_upper,omitempty"`
	VideoViewsX2sUpper     *[]int  `json:"video_views_2s_upper,omitempty"`
}
