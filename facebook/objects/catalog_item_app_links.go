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
