package objects

type AdLimitSetting struct {
	LimitAllocationByPageAdvertisers *[]map[string]int `json:"limit_allocation_by_page_advertisers,omitempty"`
}
