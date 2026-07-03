package enums

type HomelistinggetPartnerVerification string

const (
	HomelistinggetPartnerVerificationEmptyValue HomelistinggetPartnerVerification = "EMPTY_VALUE"
	HomelistinggetPartnerVerificationNone       HomelistinggetPartnerVerification = "NONE"
	HomelistinggetPartnerVerificationVerified   HomelistinggetPartnerVerification = "VERIFIED"
)

func (value HomelistinggetPartnerVerification) String() string {
	return string(value)
}
