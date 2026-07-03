package objects

type PlacementSoftOptOut struct {
	AudienceNetworkPositions *[]string `json:"audience_network_positions,omitempty"`
	FacebookPositions        *[]string `json:"facebook_positions,omitempty"`
	InstagramPositions       *[]string `json:"instagram_positions,omitempty"`
	MessengerPositions       *[]string `json:"messenger_positions,omitempty"`
	OculusPositions          *[]string `json:"oculus_positions,omitempty"`
	ThreadsPositions         *[]string `json:"threads_positions,omitempty"`
	WhatsappPositions        *[]string `json:"whatsapp_positions,omitempty"`
}

var PlacementSoftOptOutFields = struct {
	AudienceNetworkPositions string
	FacebookPositions        string
	InstagramPositions       string
	MessengerPositions       string
	OculusPositions          string
	ThreadsPositions         string
	WhatsappPositions        string
}{
	AudienceNetworkPositions: "audience_network_positions",
	FacebookPositions:        "facebook_positions",
	InstagramPositions:       "instagram_positions",
	MessengerPositions:       "messenger_positions",
	OculusPositions:          "oculus_positions",
	ThreadsPositions:         "threads_positions",
	WhatsappPositions:        "whatsapp_positions",
}
