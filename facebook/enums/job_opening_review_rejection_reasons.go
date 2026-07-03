package enums

type JobopeningReviewRejectionReasons string

const (
	JobopeningReviewRejectionReasonsAdultContent        JobopeningReviewRejectionReasons = "ADULT_CONTENT"
	JobopeningReviewRejectionReasonsDiscrimination      JobopeningReviewRejectionReasons = "DISCRIMINATION"
	JobopeningReviewRejectionReasonsDrugs               JobopeningReviewRejectionReasons = "DRUGS"
	JobopeningReviewRejectionReasonsGenericDefault      JobopeningReviewRejectionReasons = "GENERIC_DEFAULT"
	JobopeningReviewRejectionReasonsIllegal             JobopeningReviewRejectionReasons = "ILLEGAL"
	JobopeningReviewRejectionReasonsImpersonation       JobopeningReviewRejectionReasons = "IMPERSONATION"
	JobopeningReviewRejectionReasonsMisleading          JobopeningReviewRejectionReasons = "MISLEADING"
	JobopeningReviewRejectionReasonsMultilevelMarketing JobopeningReviewRejectionReasons = "MULTILEVEL_MARKETING"
	JobopeningReviewRejectionReasonsPersonalInfo        JobopeningReviewRejectionReasons = "PERSONAL_INFO"
	JobopeningReviewRejectionReasonsSexual              JobopeningReviewRejectionReasons = "SEXUAL"
)

func (value JobopeningReviewRejectionReasons) String() string {
	return string(value)
}
