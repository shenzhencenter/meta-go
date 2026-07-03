package objects

type CatalogItemAppLinks struct {
	Android          *[]AndroidAppLink      `json:"android,omitempty"`
	Ios              *[]IosAppLink          `json:"ios,omitempty"`
	Ipad             *[]IosAppLink          `json:"ipad,omitempty"`
	Iphone           *[]IosAppLink          `json:"iphone,omitempty"`
	Web              *WebAppLink            `json:"web,omitempty"`
	Windows          *[]WindowsAppLink      `json:"windows,omitempty"`
	WindowsPhone     *[]WindowsPhoneAppLink `json:"windows_phone,omitempty"`
	WindowsUniversal *[]WindowsAppLink      `json:"windows_universal,omitempty"`
}

var CatalogItemAppLinksFields = struct {
	Android          string
	Ios              string
	Ipad             string
	Iphone           string
	Web              string
	Windows          string
	WindowsPhone     string
	WindowsUniversal string
}{
	Android:          "android",
	Ios:              "ios",
	Ipad:             "ipad",
	Iphone:           "iphone",
	Web:              "web",
	Windows:          "windows",
	WindowsPhone:     "windows_phone",
	WindowsUniversal: "windows_universal",
}
