package objects

type ReachFrequencyCurveLowerConfidenceRange struct {
	ImpressionLower        *[]int  `json:"impression_lower,omitempty"`
	NumPoints              *uint64 `json:"num_points,omitempty"`
	Reach                  *[]int  `json:"reach,omitempty"`
	ReachLower             *[]int  `json:"reach_lower,omitempty"`
	UniqVideoViewsX2sLower *[]int  `json:"uniq_video_views_2s_lower,omitempty"`
	VideoViewsX2sLower     *[]int  `json:"video_views_2s_lower,omitempty"`
}
