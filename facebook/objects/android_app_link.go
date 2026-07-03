package objects

type AndroidAppLink struct {
	AppName *string `json:"app_name,omitempty"`
	Class   *string `json:"class,omitempty"`
	Package *string `json:"package,omitempty"`
	URL     *string `json:"url,omitempty"`
}
