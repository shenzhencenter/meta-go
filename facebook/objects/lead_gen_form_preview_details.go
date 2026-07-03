package objects

type LeadGenFormPreviewDetails struct {
	BookOnWebsiteText                             *string                 `json:"book_on_website_text,omitempty"`
	CallBusinessText                              *string                 `json:"call_business_text,omitempty"`
	ChatOnMessengerText                           *string                 `json:"chat_on_messenger_text,omitempty"`
	ChatOnWhatsappText                            *string                 `json:"chat_on_whatsapp_text,omitempty"`
	ContactInformationText                        *string                 `json:"contact_information_text,omitempty"`
	CreativesOverviewDefaultText                  *string                 `json:"creatives_overview_default_text,omitempty"`
	CustomDisclaimerEditorState                   *string                 `json:"custom_disclaimer_editor_state,omitempty"`
	CustomDisclaimerTitle                         *string                 `json:"custom_disclaimer_title,omitempty"`
	DataPrivacyPolicySettingDescription           *string                 `json:"data_privacy_policy_setting_description,omitempty"`
	DefaultAppointmentSchedulingInlineContext     *string                 `json:"default_appointment_scheduling_inline_context,omitempty"`
	DefaultDisqualifiedEndComponent               *map[string]interface{} `json:"default_disqualified_end_component,omitempty"`
	DefaultThankYouPage                           *map[string]interface{} `json:"default_thank_you_page,omitempty"`
	DisqualifiedThankYouCardTransparencyInfoText  *string                 `json:"disqualified_thank_you_card_transparency_info_text,omitempty"`
	EditText                                      *string                 `json:"edit_text,omitempty"`
	EmailInlineContextText                        *string                 `json:"email_inline_context_text,omitempty"`
	EmailMessengerPushOptInDisclaimer             *string                 `json:"email_messenger_push_opt_in_disclaimer,omitempty"`
	EmailMessengerPushOptInTransparencyText       *string                 `json:"email_messenger_push_opt_in_transparency_text,omitempty"`
	FormClarityDescriptionContent                 *string                 `json:"form_clarity_description_content,omitempty"`
	FormClarityDescriptionTitle                   *string                 `json:"form_clarity_description_title,omitempty"`
	FormClarityHeadline                           *string                 `json:"form_clarity_headline,omitempty"`
	GatedContentLockedDescription                 *string                 `json:"gated_content_locked_description,omitempty"`
	GatedContentLockedHeadline                    *string                 `json:"gated_content_locked_headline,omitempty"`
	GatedContentUnlockedDescription               *string                 `json:"gated_content_unlocked_description,omitempty"`
	GatedContentUnlockedHeadline                  *string                 `json:"gated_content_unlocked_headline,omitempty"`
	HowItWorksSectionHeaders                      *[]map[string]string    `json:"how_it_works_section_headers,omitempty"`
	NextButtonText                                *string                 `json:"next_button_text,omitempty"`
	OptionalQuestionText                          *string                 `json:"optional_question_text,omitempty"`
	PersonalInfoText                              *string                 `json:"personal_info_text,omitempty"`
	PhoneNumberInlineContextText                  *string                 `json:"phone_number_inline_context_text,omitempty"`
	PrivacyPolicyLinkText                         *string                 `json:"privacy_policy_link_text,omitempty"`
	PrivacyPolicyLinkTextForOptionalPrivacyPolicy *string                 `json:"privacy_policy_link_text_for_optional_privacy_policy,omitempty"`
	PrivacyPolicyTitleSectionTitleText            *string                 `json:"privacy_policy_title_section_title_text,omitempty"`
	PrivacySettingDescription                     *string                 `json:"privacy_setting_description,omitempty"`
	ProductsSectionHeaders                        *[]map[string]string    `json:"products_section_headers,omitempty"`
	QualifiedThankYouCardTransparencyInfoText     *string                 `json:"qualified_thank_you_card_transparency_info_text,omitempty"`
	RedeemPromoCodeText                           *string                 `json:"redeem_promo_code_text,omitempty"`
	ReturnToFacebookText                          *string                 `json:"return_to_facebook_text,omitempty"`
	ReviewYourInfoText                            *string                 `json:"review_your_info_text,omitempty"`
	SecureSharingText                             *string                 `json:"secure_sharing_text,omitempty"`
	SecureSharingTextForEmbeddedBookingCalendly   *string                 `json:"secure_sharing_text_for_embedded_booking_calendly,omitempty"`
	SecureSharingTextForEmbeddedBookingGhl        *string                 `json:"secure_sharing_text_for_embedded_booking_ghl,omitempty"`
	SecureSharingTextForOptionalPrivacyPolicy     *string                 `json:"secure_sharing_text_for_optional_privacy_policy,omitempty"`
	SlideToSubmitText                             *string                 `json:"slide_to_submit_text,omitempty"`
	SocialProofSectionHeaders                     *[]map[string]string    `json:"social_proof_section_headers,omitempty"`
	SubmitButtonText                              *string                 `json:"submit_button_text,omitempty"`
	ViewFileText                                  *string                 `json:"view_file_text,omitempty"`
	WhatsAppOptInBody                             *string                 `json:"whats_app_opt_in_body,omitempty"`
	WhatsAppOptInTitle                            *string                 `json:"whats_app_opt_in_title,omitempty"`
}

var LeadGenFormPreviewDetailsFields = struct {
	BookOnWebsiteText                             string
	CallBusinessText                              string
	ChatOnMessengerText                           string
	ChatOnWhatsappText                            string
	ContactInformationText                        string
	CreativesOverviewDefaultText                  string
	CustomDisclaimerEditorState                   string
	CustomDisclaimerTitle                         string
	DataPrivacyPolicySettingDescription           string
	DefaultAppointmentSchedulingInlineContext     string
	DefaultDisqualifiedEndComponent               string
	DefaultThankYouPage                           string
	DisqualifiedThankYouCardTransparencyInfoText  string
	EditText                                      string
	EmailInlineContextText                        string
	EmailMessengerPushOptInDisclaimer             string
	EmailMessengerPushOptInTransparencyText       string
	FormClarityDescriptionContent                 string
	FormClarityDescriptionTitle                   string
	FormClarityHeadline                           string
	GatedContentLockedDescription                 string
	GatedContentLockedHeadline                    string
	GatedContentUnlockedDescription               string
	GatedContentUnlockedHeadline                  string
	HowItWorksSectionHeaders                      string
	NextButtonText                                string
	OptionalQuestionText                          string
	PersonalInfoText                              string
	PhoneNumberInlineContextText                  string
	PrivacyPolicyLinkText                         string
	PrivacyPolicyLinkTextForOptionalPrivacyPolicy string
	PrivacyPolicyTitleSectionTitleText            string
	PrivacySettingDescription                     string
	ProductsSectionHeaders                        string
	QualifiedThankYouCardTransparencyInfoText     string
	RedeemPromoCodeText                           string
	ReturnToFacebookText                          string
	ReviewYourInfoText                            string
	SecureSharingText                             string
	SecureSharingTextForEmbeddedBookingCalendly   string
	SecureSharingTextForEmbeddedBookingGhl        string
	SecureSharingTextForOptionalPrivacyPolicy     string
	SlideToSubmitText                             string
	SocialProofSectionHeaders                     string
	SubmitButtonText                              string
	ViewFileText                                  string
	WhatsAppOptInBody                             string
	WhatsAppOptInTitle                            string
}{
	BookOnWebsiteText:                            "book_on_website_text",
	CallBusinessText:                             "call_business_text",
	ChatOnMessengerText:                          "chat_on_messenger_text",
	ChatOnWhatsappText:                           "chat_on_whatsapp_text",
	ContactInformationText:                       "contact_information_text",
	CreativesOverviewDefaultText:                 "creatives_overview_default_text",
	CustomDisclaimerEditorState:                  "custom_disclaimer_editor_state",
	CustomDisclaimerTitle:                        "custom_disclaimer_title",
	DataPrivacyPolicySettingDescription:          "data_privacy_policy_setting_description",
	DefaultAppointmentSchedulingInlineContext:    "default_appointment_scheduling_inline_context",
	DefaultDisqualifiedEndComponent:              "default_disqualified_end_component",
	DefaultThankYouPage:                          "default_thank_you_page",
	DisqualifiedThankYouCardTransparencyInfoText: "disqualified_thank_you_card_transparency_info_text",
	EditText:                                      "edit_text",
	EmailInlineContextText:                        "email_inline_context_text",
	EmailMessengerPushOptInDisclaimer:             "email_messenger_push_opt_in_disclaimer",
	EmailMessengerPushOptInTransparencyText:       "email_messenger_push_opt_in_transparency_text",
	FormClarityDescriptionContent:                 "form_clarity_description_content",
	FormClarityDescriptionTitle:                   "form_clarity_description_title",
	FormClarityHeadline:                           "form_clarity_headline",
	GatedContentLockedDescription:                 "gated_content_locked_description",
	GatedContentLockedHeadline:                    "gated_content_locked_headline",
	GatedContentUnlockedDescription:               "gated_content_unlocked_description",
	GatedContentUnlockedHeadline:                  "gated_content_unlocked_headline",
	HowItWorksSectionHeaders:                      "how_it_works_section_headers",
	NextButtonText:                                "next_button_text",
	OptionalQuestionText:                          "optional_question_text",
	PersonalInfoText:                              "personal_info_text",
	PhoneNumberInlineContextText:                  "phone_number_inline_context_text",
	PrivacyPolicyLinkText:                         "privacy_policy_link_text",
	PrivacyPolicyLinkTextForOptionalPrivacyPolicy: "privacy_policy_link_text_for_optional_privacy_policy",
	PrivacyPolicyTitleSectionTitleText:            "privacy_policy_title_section_title_text",
	PrivacySettingDescription:                     "privacy_setting_description",
	ProductsSectionHeaders:                        "products_section_headers",
	QualifiedThankYouCardTransparencyInfoText:     "qualified_thank_you_card_transparency_info_text",
	RedeemPromoCodeText:                           "redeem_promo_code_text",
	ReturnToFacebookText:                          "return_to_facebook_text",
	ReviewYourInfoText:                            "review_your_info_text",
	SecureSharingText:                             "secure_sharing_text",
	SecureSharingTextForEmbeddedBookingCalendly:   "secure_sharing_text_for_embedded_booking_calendly",
	SecureSharingTextForEmbeddedBookingGhl:        "secure_sharing_text_for_embedded_booking_ghl",
	SecureSharingTextForOptionalPrivacyPolicy:     "secure_sharing_text_for_optional_privacy_policy",
	SlideToSubmitText:                             "slide_to_submit_text",
	SocialProofSectionHeaders:                     "social_proof_section_headers",
	SubmitButtonText:                              "submit_button_text",
	ViewFileText:                                  "view_file_text",
	WhatsAppOptInBody:                             "whats_app_opt_in_body",
	WhatsAppOptInTitle:                            "whats_app_opt_in_title",
}
