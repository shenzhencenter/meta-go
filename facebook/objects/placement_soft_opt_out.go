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
