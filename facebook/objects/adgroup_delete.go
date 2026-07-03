package objects

type AdgroupDelete struct {
	Success *bool `json:"success,omitempty"`
}

var AdgroupDeleteFields = struct {
	Success string
}{
	Success: "success",
}
