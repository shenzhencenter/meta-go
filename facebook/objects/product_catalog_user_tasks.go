package objects

type ProductCatalogUserTasks struct {
	Business *Business `json:"business,omitempty"`
	Tasks    *[]string `json:"tasks,omitempty"`
}
