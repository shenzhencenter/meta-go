package objects

type CustomUserSettings struct {
	PageLevelPersistentMenu *[]map[string]interface{} `json:"page_level_persistent_menu,omitempty"`
	UserLevelPersistentMenu *[]map[string]interface{} `json:"user_level_persistent_menu,omitempty"`
}
