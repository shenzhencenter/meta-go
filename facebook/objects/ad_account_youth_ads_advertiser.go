package objects

type AdAccountYouthAdsAdvertiser struct {
	IsYouthAdsAdvertiser *bool `json:"is_youth_ads_advertiser,omitempty"`
}

var AdAccountYouthAdsAdvertiserFields = struct {
	IsYouthAdsAdvertiser string
}{
	IsYouthAdsAdvertiser: "is_youth_ads_advertiser",
}
