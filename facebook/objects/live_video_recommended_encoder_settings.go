package objects

type LiveVideoRecommendedEncoderSettings struct {
	AudioCodecSettings *map[string]interface{} `json:"audio_codec_settings,omitempty"`
	StreamingProtocol  *string                 `json:"streaming_protocol,omitempty"`
	VideoCodecSettings *map[string]interface{} `json:"video_codec_settings,omitempty"`
}
