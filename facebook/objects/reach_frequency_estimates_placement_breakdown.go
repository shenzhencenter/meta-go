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

var ReachFrequencyEstimatesPlacementBreakdownFields = struct {
	Android         string
	AudienceNetwork string
	Desktop         string
	FacebookSearch  string
	FbReels         string
	FbReelsOverlay  string
	IgAndroid       string
	IgIos           string
	IgOther         string
	IgReels         string
	IgStory         string
	InstantArticles string
	InstreamVideos  string
	Ios             string
	Msite           string
	SuggestedVideos string
}{
	Android:         "android",
	AudienceNetwork: "audience_network",
	Desktop:         "desktop",
	FacebookSearch:  "facebook_search",
	FbReels:         "fb_reels",
	FbReelsOverlay:  "fb_reels_overlay",
	IgAndroid:       "ig_android",
	IgIos:           "ig_ios",
	IgOther:         "ig_other",
	IgReels:         "ig_reels",
	IgStory:         "ig_story",
	InstantArticles: "instant_articles",
	InstreamVideos:  "instream_videos",
	Ios:             "ios",
	Msite:           "msite",
	SuggestedVideos: "suggested_videos",
}
