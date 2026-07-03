package objects

type ReachFrequencyEstimatesPlacementBreakdown struct {
	Android         *[]float64 `json:"android,omitempty"`
	AudienceNetwork *[]float64 `json:"audience_network,omitempty"`
	Desktop         *[]float64 `json:"desktop,omitempty"`
	FacebookSearch  *[]float64 `json:"facebook_search,omitempty"`
	FbReels         *[]float64 `json:"fb_reels,omitempty"`
	FbReelsOverlay  *[]float64 `json:"fb_reels_overlay,omitempty"`
	IgAndroid       *[]float64 `json:"ig_android,omitempty"`
	IgIos           *[]float64 `json:"ig_ios,omitempty"`
	IgOther         *[]float64 `json:"ig_other,omitempty"`
	IgReels         *[]float64 `json:"ig_reels,omitempty"`
	IgStory         *[]float64 `json:"ig_story,omitempty"`
	InstantArticles *[]float64 `json:"instant_articles,omitempty"`
	InstreamVideos  *[]float64 `json:"instream_videos,omitempty"`
	Ios             *[]float64 `json:"ios,omitempty"`
	Msite           *[]float64 `json:"msite,omitempty"`
	SuggestedVideos *[]float64 `json:"suggested_videos,omitempty"`
}
