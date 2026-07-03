package enums

type ApplicationwhatsappBusinessSolutionsRoleEnumParam string

const (
	ApplicationwhatsappBusinessSolutionsRoleEnumParamOwner   ApplicationwhatsappBusinessSolutionsRoleEnumParam = "OWNER"
	ApplicationwhatsappBusinessSolutionsRoleEnumParamPartner ApplicationwhatsappBusinessSolutionsRoleEnumParam = "PARTNER"
)

func (value ApplicationwhatsappBusinessSolutionsRoleEnumParam) String() string {
	return string(value)
}
