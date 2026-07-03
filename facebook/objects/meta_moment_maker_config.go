package objects

type MetaMomentMakerConfig struct {
	SaturationMode *string `json:"saturation_mode,omitempty"`
}

var MetaMomentMakerConfigFields = struct {
	SaturationMode string
}{
	SaturationMode: "saturation_mode",
}
