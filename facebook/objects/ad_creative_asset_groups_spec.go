package objects

type AdCreativeAssetGroupsSpec struct {
	Groups  *[]map[string]interface{} `json:"groups,omitempty"`
	Origin  *string                   `json:"origin,omitempty"`
	Origins *[]string                 `json:"origins,omitempty"`
}
