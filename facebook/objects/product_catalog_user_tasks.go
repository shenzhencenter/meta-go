package objects

type ProductCatalogUserTasks struct {
	Business *Business `json:"business,omitempty"`
	Tasks    *[]string `json:"tasks,omitempty"`
}

var ProductCatalogUserTasksFields = struct {
	Business string
	Tasks    string
}{
	Business: "business",
	Tasks:    "tasks",
}
