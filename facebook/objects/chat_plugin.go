package objects

type ChatPlugin struct {
	Alignment             *string `json:"alignment,omitempty"`
	DesktopBottomSpacing  *string `json:"desktop_bottom_spacing,omitempty"`
	DesktopSideSpacing    *string `json:"desktop_side_spacing,omitempty"`
	EntryPointIcon        *string `json:"entry_point_icon,omitempty"`
	EntryPointLabel       *string `json:"entry_point_label,omitempty"`
	GreetingDialogDisplay *string `json:"greeting_dialog_display,omitempty"`
	GuestChatMode         *string `json:"guest_chat_mode,omitempty"`
	MobileBottomSpacing   *string `json:"mobile_bottom_spacing,omitempty"`
	MobileChatDisplay     *string `json:"mobile_chat_display,omitempty"`
	MobileSideSpacing     *string `json:"mobile_side_spacing,omitempty"`
	ThemeColor            *string `json:"theme_color,omitempty"`
	WelcomeScreenGreeting *string `json:"welcome_screen_greeting,omitempty"`
}

var ChatPluginFields = struct {
	Alignment             string
	DesktopBottomSpacing  string
	DesktopSideSpacing    string
	EntryPointIcon        string
	EntryPointLabel       string
	GreetingDialogDisplay string
	GuestChatMode         string
	MobileBottomSpacing   string
	MobileChatDisplay     string
	MobileSideSpacing     string
	ThemeColor            string
	WelcomeScreenGreeting string
}{
	Alignment:             "alignment",
	DesktopBottomSpacing:  "desktop_bottom_spacing",
	DesktopSideSpacing:    "desktop_side_spacing",
	EntryPointIcon:        "entry_point_icon",
	EntryPointLabel:       "entry_point_label",
	GreetingDialogDisplay: "greeting_dialog_display",
	GuestChatMode:         "guest_chat_mode",
	MobileBottomSpacing:   "mobile_bottom_spacing",
	MobileChatDisplay:     "mobile_chat_display",
	MobileSideSpacing:     "mobile_side_spacing",
	ThemeColor:            "theme_color",
	WelcomeScreenGreeting: "welcome_screen_greeting",
}
