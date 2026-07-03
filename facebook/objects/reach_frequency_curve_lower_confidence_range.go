package objects

type ReachFrequencyCurveLowerConfidenceRange struct {
	ImpressionLower        *[]int  `json:"impression_lower,omitempty"`
	NumPoints              *uint64 `json:"num_points,omitempty"`
	Reach                  *[]int  `json:"reach,omitempty"`
	ReachLower             *[]int  `json:"reach_lower,omitempty"`
	UniqVideoViewsX2sLower *[]int  `json:"uniq_video_views_2s_lower,omitempty"`
	VideoViewsX2sLower     *[]int  `json:"video_views_2s_lower,omitempty"`
}

var ReachFrequencyCurveLowerConfidenceRangeFields = struct {
	ImpressionLower        string
	NumPoints              string
	Reach                  string
	ReachLower             string
	UniqVideoViewsX2sLower string
	VideoViewsX2sLower     string
}{
	ImpressionLower:        "impression_lower",
	NumPoints:              "num_points",
	Reach:                  "reach",
	ReachLower:             "reach_lower",
	UniqVideoViewsX2sLower: "uniq_video_views_2s_lower",
	VideoViewsX2sLower:     "video_views_2s_lower",
}
