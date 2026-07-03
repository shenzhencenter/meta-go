package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AppLinks struct {
	Android          *[]AndroidAppLink      `json:"android,omitempty"`
	ID               *core.ID               `json:"id,omitempty"`
	Ios              *[]IosAppLink          `json:"ios,omitempty"`
	Ipad             *[]IosAppLink          `json:"ipad,omitempty"`
	Iphone           *[]IosAppLink          `json:"iphone,omitempty"`
	Web              *WebAppLink            `json:"web,omitempty"`
	Windows          *[]WindowsAppLink      `json:"windows,omitempty"`
	WindowsPhone     *[]WindowsPhoneAppLink `json:"windows_phone,omitempty"`
	WindowsUniversal *[]WindowsAppLink      `json:"windows_universal,omitempty"`
}

var AppLinksFields = struct {
	Android          string
	ID               string
	Ios              string
	Ipad             string
	Iphone           string
	Web              string
	Windows          string
	WindowsPhone     string
	WindowsUniversal string
}{
	Android:          "android",
	ID:               "id",
	Ios:              "ios",
	Ipad:             "ipad",
	Iphone:           "iphone",
	Web:              "web",
	Windows:          "windows",
	WindowsPhone:     "windows_phone",
	WindowsUniversal: "windows_universal",
}
