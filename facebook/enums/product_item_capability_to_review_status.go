package enums

type ProductitemCapabilityToReviewStatus string

const (
	ProductitemCapabilityToReviewStatusApproved ProductitemCapabilityToReviewStatus = "APPROVED"
	ProductitemCapabilityToReviewStatusNoReview ProductitemCapabilityToReviewStatus = "NO_REVIEW"
	ProductitemCapabilityToReviewStatusOutdated ProductitemCapabilityToReviewStatus = "OUTDATED"
	ProductitemCapabilityToReviewStatusPending  ProductitemCapabilityToReviewStatus = "PENDING"
	ProductitemCapabilityToReviewStatusRejected ProductitemCapabilityToReviewStatus = "REJECTED"
)

func (value ProductitemCapabilityToReviewStatus) String() string {
	return string(value)
}
