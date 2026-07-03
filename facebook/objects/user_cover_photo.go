package objects

type UserCoverPhoto struct {
	OffsetX *float64 `json:"offset_x,omitempty"`
	OffsetY *float64 `json:"offset_y,omitempty"`
	Source  *string  `json:"source,omitempty"`
}

var UserCoverPhotoFields = struct {
	OffsetX string
	OffsetY string
	Source  string
}{
	OffsetX: "offset_x",
	OffsetY: "offset_y",
	Source:  "source",
}
