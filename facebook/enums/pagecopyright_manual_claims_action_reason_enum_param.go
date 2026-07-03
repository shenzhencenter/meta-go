package enums

type PagecopyrightManualClaimsActionReasonEnumParam string

const (
	PagecopyrightManualClaimsActionReasonEnumParamArticleX17Preflagging     PagecopyrightManualClaimsActionReasonEnumParam = "ARTICLE_17_PREFLAGGING"
	PagecopyrightManualClaimsActionReasonEnumParamArtistObjection           PagecopyrightManualClaimsActionReasonEnumParam = "ARTIST_OBJECTION"
	PagecopyrightManualClaimsActionReasonEnumParamObjectionableContent      PagecopyrightManualClaimsActionReasonEnumParam = "OBJECTIONABLE_CONTENT"
	PagecopyrightManualClaimsActionReasonEnumParamPremiumMusicVideo         PagecopyrightManualClaimsActionReasonEnumParam = "PREMIUM_MUSIC_VIDEO"
	PagecopyrightManualClaimsActionReasonEnumParamPrereleaseContent         PagecopyrightManualClaimsActionReasonEnumParam = "PRERELEASE_CONTENT"
	PagecopyrightManualClaimsActionReasonEnumParamProductParameters         PagecopyrightManualClaimsActionReasonEnumParam = "PRODUCT_PARAMETERS"
	PagecopyrightManualClaimsActionReasonEnumParamRestrictedContent         PagecopyrightManualClaimsActionReasonEnumParam = "RESTRICTED_CONTENT"
	PagecopyrightManualClaimsActionReasonEnumParamUnauthorizedCommercialUse PagecopyrightManualClaimsActionReasonEnumParam = "UNAUTHORIZED_COMMERCIAL_USE"
)

func (value PagecopyrightManualClaimsActionReasonEnumParam) String() string {
	return string(value)
}
