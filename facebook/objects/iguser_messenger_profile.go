package objects

type IGUserMessengerProfile struct {
	IceBreakers    *[]map[string]interface{} `json:"ice_breakers,omitempty"`
	PersistentMenu *[]map[string]interface{} `json:"persistent_menu,omitempty"`
}
