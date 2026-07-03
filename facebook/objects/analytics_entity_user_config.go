package objects

type AnalyticsEntityUserConfig struct {
	DismissedNotices *[]string `json:"dismissed_notices,omitempty"`
}

var AnalyticsEntityUserConfigFields = struct {
	DismissedNotices string
}{
	DismissedNotices: "dismissed_notices",
}
