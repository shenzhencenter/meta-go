package objects

type AdCreativeTemplateURLSpec struct {
	Android      *map[string]interface{} `json:"android,omitempty"`
	Config       *map[string]interface{} `json:"config,omitempty"`
	Ios          *map[string]interface{} `json:"ios,omitempty"`
	Ipad         *map[string]interface{} `json:"ipad,omitempty"`
	Iphone       *map[string]interface{} `json:"iphone,omitempty"`
	Web          *map[string]interface{} `json:"web,omitempty"`
	WindowsPhone *map[string]interface{} `json:"windows_phone,omitempty"`
}

var AdCreativeTemplateURLSpecFields = struct {
	Android      string
	Config       string
	Ios          string
	Ipad         string
	Iphone       string
	Web          string
	WindowsPhone string
}{
	Android:      "android",
	Config:       "config",
	Ios:          "ios",
	Ipad:         "ipad",
	Iphone:       "iphone",
	Web:          "web",
	WindowsPhone: "windows_phone",
}
