package objects

type AdCreativeFormatTransformationSpec struct {
	Customizations *[]map[string]interface{} `json:"customizations,omitempty"`
	DataSource     *[]string                 `json:"data_source,omitempty"`
	Format         *string                   `json:"format,omitempty"`
}
