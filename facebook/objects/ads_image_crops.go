package objects

type AdsImageCrops struct {
	X100x100 *[][]interface{} `json:"100x100,omitempty"`
	X100x72  *[][]interface{} `json:"100x72,omitempty"`
	X191x100 *[][]interface{} `json:"191x100,omitempty"`
	X300x400 *[][]interface{} `json:"300x400,omitempty"`
	X400x150 *[][]interface{} `json:"400x150,omitempty"`
	X400x500 *[][]interface{} `json:"400x500,omitempty"`
	X600x360 *[][]interface{} `json:"600x360,omitempty"`
	X90x160  *[][]interface{} `json:"90x160,omitempty"`
}

var AdsImageCropsFields = struct {
	X100x100 string
	X100x72  string
	X191x100 string
	X300x400 string
	X400x150 string
	X400x500 string
	X600x360 string
	X90x160  string
}{
	X100x100: "100x100",
	X100x72:  "100x72",
	X191x100: "191x100",
	X300x400: "300x400",
	X400x150: "400x150",
	X400x500: "400x500",
	X600x360: "600x360",
	X90x160:  "90x160",
}
