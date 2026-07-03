package enums

type WhatsappbusinessaccountassignedUsersTasksEnumParam string

const (
	WhatsappbusinessaccountassignedUsersTasksEnumParamDevelop           WhatsappbusinessaccountassignedUsersTasksEnumParam = "DEVELOP"
	WhatsappbusinessaccountassignedUsersTasksEnumParamManage            WhatsappbusinessaccountassignedUsersTasksEnumParam = "MANAGE"
	WhatsappbusinessaccountassignedUsersTasksEnumParamManageExtensions  WhatsappbusinessaccountassignedUsersTasksEnumParam = "MANAGE_EXTENSIONS"
	WhatsappbusinessaccountassignedUsersTasksEnumParamManagePhone       WhatsappbusinessaccountassignedUsersTasksEnumParam = "MANAGE_PHONE"
	WhatsappbusinessaccountassignedUsersTasksEnumParamManagePhoneAssets WhatsappbusinessaccountassignedUsersTasksEnumParam = "MANAGE_PHONE_ASSETS"
	WhatsappbusinessaccountassignedUsersTasksEnumParamManageTemplates   WhatsappbusinessaccountassignedUsersTasksEnumParam = "MANAGE_TEMPLATES"
	WhatsappbusinessaccountassignedUsersTasksEnumParamMessaging         WhatsappbusinessaccountassignedUsersTasksEnumParam = "MESSAGING"
	WhatsappbusinessaccountassignedUsersTasksEnumParamViewCost          WhatsappbusinessaccountassignedUsersTasksEnumParam = "VIEW_COST"
	WhatsappbusinessaccountassignedUsersTasksEnumParamViewPhoneAssets   WhatsappbusinessaccountassignedUsersTasksEnumParam = "VIEW_PHONE_ASSETS"
	WhatsappbusinessaccountassignedUsersTasksEnumParamViewTemplates     WhatsappbusinessaccountassignedUsersTasksEnumParam = "VIEW_TEMPLATES"
)

func (value WhatsappbusinessaccountassignedUsersTasksEnumParam) String() string {
	return string(value)
}
