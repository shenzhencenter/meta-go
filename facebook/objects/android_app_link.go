package objects

type AndroidAppLink struct {
	AppName *string `json:"app_name,omitempty"`
	Class   *string `json:"class,omitempty"`
	Package *string `json:"package,omitempty"`
	URL     *string `json:"url,omitempty"`
}

var AndroidAppLinkFields = struct {
	AppName string
	Class   string
	Package string
	URL     string
}{
	AppName: "app_name",
	Class:   "class",
	Package: "package",
	URL:     "url",
}
