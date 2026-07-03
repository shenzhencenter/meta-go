package enums

type ProductitemReviewStatus string

const (
	ProductitemReviewStatusValue    ProductitemReviewStatus = ""
	ProductitemReviewStatusApproved ProductitemReviewStatus = "approved"
	ProductitemReviewStatusOutdated ProductitemReviewStatus = "outdated"
	ProductitemReviewStatusPending  ProductitemReviewStatus = "pending"
	ProductitemReviewStatusRejected ProductitemReviewStatus = "rejected"
)

func (value ProductitemReviewStatus) String() string {
	return string(value)
}
