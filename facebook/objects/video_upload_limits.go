package objects

type VideoUploadLimits struct {
	Length *uint64 `json:"length,omitempty"`
	Size   *int    `json:"size,omitempty"`
}
