package enums

type UserLocalNewsSubscriptionStatus string

const (
	UserLocalNewsSubscriptionStatusStatusOff UserLocalNewsSubscriptionStatus = "STATUS_OFF"
	UserLocalNewsSubscriptionStatusStatusOn  UserLocalNewsSubscriptionStatus = "STATUS_ON"
)

func (value UserLocalNewsSubscriptionStatus) String() string {
	return string(value)
}
