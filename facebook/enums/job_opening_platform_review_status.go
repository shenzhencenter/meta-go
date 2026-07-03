package enums

type JobopeningPlatformReviewStatus string

const (
	JobopeningPlatformReviewStatusApproved JobopeningPlatformReviewStatus = "APPROVED"
	JobopeningPlatformReviewStatusPending  JobopeningPlatformReviewStatus = "PENDING"
	JobopeningPlatformReviewStatusRejected JobopeningPlatformReviewStatus = "REJECTED"
)

func (value JobopeningPlatformReviewStatus) String() string {
	return string(value)
}
