package enums

type ThirdpartypartnerviewabilityrequestPlatform string

const (
	ThirdpartypartnerviewabilityrequestPlatformAudienceNetwork ThirdpartypartnerviewabilityrequestPlatform = "AUDIENCE_NETWORK"
	ThirdpartypartnerviewabilityrequestPlatformFacebook        ThirdpartypartnerviewabilityrequestPlatform = "FACEBOOK"
	ThirdpartypartnerviewabilityrequestPlatformInstagram       ThirdpartypartnerviewabilityrequestPlatform = "INSTAGRAM"
)

func (value ThirdpartypartnerviewabilityrequestPlatform) String() string {
	return string(value)
}
