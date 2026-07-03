package enums

type UserLocalNewsMegaphoneDismissStatus string

const (
	UserLocalNewsMegaphoneDismissStatusNo  UserLocalNewsMegaphoneDismissStatus = "NO"
	UserLocalNewsMegaphoneDismissStatusYes UserLocalNewsMegaphoneDismissStatus = "YES"
)

func (value UserLocalNewsMegaphoneDismissStatus) String() string {
	return string(value)
}
