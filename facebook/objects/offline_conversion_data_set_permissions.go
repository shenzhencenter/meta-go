package objects

type OfflineConversionDataSetPermissions struct {
	CanEdit                                  *bool `json:"can_edit,omitempty"`
	CanEditOrUpload                          *bool `json:"can_edit_or_upload,omitempty"`
	CanUpload                                *bool `json:"can_upload,omitempty"`
	ShouldBlockVanillaBusinessEmployeeAccess *bool `json:"should_block_vanilla_business_employee_access,omitempty"`
}
