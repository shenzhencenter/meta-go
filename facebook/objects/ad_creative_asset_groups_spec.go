package objects

type AdCreativeAssetGroupsSpec struct {
	Groups  *[]map[string]interface{} `json:"groups,omitempty"`
	Origin  *string                   `json:"origin,omitempty"`
	Origins *[]string                 `json:"origins,omitempty"`
}

var AdCreativeAssetGroupsSpecFields = struct {
	Groups  string
	Origin  string
	Origins string
}{
	Groups:  "groups",
	Origin:  "origin",
	Origins: "origins",
}
