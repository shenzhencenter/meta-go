package objects

type AdAccountAdsRecommendedAudios struct {
	AudioAssets *[]int `json:"audio_assets,omitempty"`
}

var AdAccountAdsRecommendedAudiosFields = struct {
	AudioAssets string
}{
	AudioAssets: "audio_assets",
}
