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
