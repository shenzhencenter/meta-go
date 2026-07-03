package objects

type ReachFrequencyCurveUpperConfidenceRange struct {
	ImpressionUpper        *[]int  `json:"impression_upper,omitempty"`
	NumPoints              *uint64 `json:"num_points,omitempty"`
	Reach                  *[]int  `json:"reach,omitempty"`
	ReachUpper             *[]int  `json:"reach_upper,omitempty"`
	UniqVideoViewsX2sUpper *[]int  `json:"uniq_video_views_2s_upper,omitempty"`
	VideoViewsX2sUpper     *[]int  `json:"video_views_2s_upper,omitempty"`
}

var ReachFrequencyCurveUpperConfidenceRangeFields = struct {
	ImpressionUpper        string
	NumPoints              string
	Reach                  string
	ReachUpper             string
	UniqVideoViewsX2sUpper string
	VideoViewsX2sUpper     string
}{
	ImpressionUpper:        "impression_upper",
	NumPoints:              "num_points",
	Reach:                  "reach",
	ReachUpper:             "reach_upper",
	UniqVideoViewsX2sUpper: "uniq_video_views_2s_upper",
	VideoViewsX2sUpper:     "video_views_2s_upper",
}
