package objects

type UserCoverPhoto struct {
	OffsetX *float64 `json:"offset_x,omitempty"`
	OffsetY *float64 `json:"offset_y,omitempty"`
	Source  *string  `json:"source,omitempty"`
}
