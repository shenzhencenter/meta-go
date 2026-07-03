package enums

type BusinessTwoFactorType string

const (
	BusinessTwoFactorTypeAdminRequired BusinessTwoFactorType = "admin_required"
	BusinessTwoFactorTypeAllRequired   BusinessTwoFactorType = "all_required"
	BusinessTwoFactorTypeNone          BusinessTwoFactorType = "none"
)

func (value BusinessTwoFactorType) String() string {
	return string(value)
}
