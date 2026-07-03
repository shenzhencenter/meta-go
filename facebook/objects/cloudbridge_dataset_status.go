package objects

type CloudbridgeDatasetStatus struct {
	AppRedactedEvent   *[]string              `json:"app_redacted_event,omitempty"`
	AppSensitiveParams *[]map[string][]string `json:"app_sensitive_params,omitempty"`
	AppUnverifiedEvent *[]string              `json:"app_unverified_event,omitempty"`
	HasAppAssociated   *bool                  `json:"has_app_associated,omitempty"`
	IsAppProhibited    *bool                  `json:"is_app_prohibited,omitempty"`
	IsDataset          *bool                  `json:"is_dataset,omitempty"`
}
