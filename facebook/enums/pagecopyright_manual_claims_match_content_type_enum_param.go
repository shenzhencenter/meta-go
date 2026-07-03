package enums

type PagecopyrightManualClaimsMatchContentTypeEnumParam string

const (
	PagecopyrightManualClaimsMatchContentTypeEnumParamAudioOnly     PagecopyrightManualClaimsMatchContentTypeEnumParam = "AUDIO_ONLY"
	PagecopyrightManualClaimsMatchContentTypeEnumParamVideoAndAudio PagecopyrightManualClaimsMatchContentTypeEnumParam = "VIDEO_AND_AUDIO"
	PagecopyrightManualClaimsMatchContentTypeEnumParamVideoOnly     PagecopyrightManualClaimsMatchContentTypeEnumParam = "VIDEO_ONLY"
)

func (value PagecopyrightManualClaimsMatchContentTypeEnumParam) String() string {
	return string(value)
}
