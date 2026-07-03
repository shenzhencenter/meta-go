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
