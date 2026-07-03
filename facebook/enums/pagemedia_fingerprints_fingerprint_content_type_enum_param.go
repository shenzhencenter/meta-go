package enums

type PagemediaFingerprintsFingerprintContentTypeEnumParam string

const (
	PagemediaFingerprintsFingerprintContentTypeEnumParamAmSongtrack PagemediaFingerprintsFingerprintContentTypeEnumParam = "AM_SONGTRACK"
	PagemediaFingerprintsFingerprintContentTypeEnumParamEpisode     PagemediaFingerprintsFingerprintContentTypeEnumParam = "EPISODE"
	PagemediaFingerprintsFingerprintContentTypeEnumParamMovie       PagemediaFingerprintsFingerprintContentTypeEnumParam = "MOVIE"
	PagemediaFingerprintsFingerprintContentTypeEnumParamOther       PagemediaFingerprintsFingerprintContentTypeEnumParam = "OTHER"
	PagemediaFingerprintsFingerprintContentTypeEnumParamSongtrack   PagemediaFingerprintsFingerprintContentTypeEnumParam = "SONGTRACK"
)

func (value PagemediaFingerprintsFingerprintContentTypeEnumParam) String() string {
	return string(value)
}
