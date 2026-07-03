package objects

type TargetingProspectingAudience struct {
	Sources *[]map[string]interface{} `json:"sources,omitempty"`
}

var TargetingProspectingAudienceFields = struct {
	Sources string
}{
	Sources: "sources",
}
