package objects

type AdCreativeLinkDataAppLinkSpec struct {
	Android *[]AndroidAppLink `json:"android,omitempty"`
	Ios     *[]IosAppLink     `json:"ios,omitempty"`
	Ipad    *[]IosAppLink     `json:"ipad,omitempty"`
	Iphone  *[]IosAppLink     `json:"iphone,omitempty"`
}

var AdCreativeLinkDataAppLinkSpecFields = struct {
	Android string
	Ios     string
	Ipad    string
	Iphone  string
}{
	Android: "android",
	Ios:     "ios",
	Ipad:    "ipad",
	Iphone:  "iphone",
}
