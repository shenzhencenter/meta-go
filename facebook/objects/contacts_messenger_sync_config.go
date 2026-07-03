package objects

type ContactsMessengerSyncConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
}

var ContactsMessengerSyncConfigFields = struct {
	Enabled string
}{
	Enabled: "enabled",
}
