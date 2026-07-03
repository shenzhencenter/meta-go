package enums

type WhatsappbusinesspartnerclientverificationsubmissionRejectionReasons string

const (
	WhatsappbusinesspartnerclientverificationsubmissionRejectionReasonsAddressNotMatching           WhatsappbusinesspartnerclientverificationsubmissionRejectionReasons = "ADDRESS_NOT_MATCHING"
	WhatsappbusinesspartnerclientverificationsubmissionRejectionReasonsBusinessNotEligible          WhatsappbusinesspartnerclientverificationsubmissionRejectionReasons = "BUSINESS_NOT_ELIGIBLE"
	WhatsappbusinesspartnerclientverificationsubmissionRejectionReasonsLegalNameNotFoundInDocuments WhatsappbusinesspartnerclientverificationsubmissionRejectionReasons = "LEGAL_NAME_NOT_FOUND_IN_DOCUMENTS"
	WhatsappbusinesspartnerclientverificationsubmissionRejectionReasonsLegalNameNotMatching         WhatsappbusinesspartnerclientverificationsubmissionRejectionReasons = "LEGAL_NAME_NOT_MATCHING"
	WhatsappbusinesspartnerclientverificationsubmissionRejectionReasonsMalformedDocuments           WhatsappbusinesspartnerclientverificationsubmissionRejectionReasons = "MALFORMED_DOCUMENTS"
	WhatsappbusinesspartnerclientverificationsubmissionRejectionReasonsNone                         WhatsappbusinesspartnerclientverificationsubmissionRejectionReasons = "NONE"
	WhatsappbusinesspartnerclientverificationsubmissionRejectionReasonsWebsiteNotMatching           WhatsappbusinesspartnerclientverificationsubmissionRejectionReasons = "WEBSITE_NOT_MATCHING"
)

func (value WhatsappbusinesspartnerclientverificationsubmissionRejectionReasons) String() string {
	return string(value)
}
