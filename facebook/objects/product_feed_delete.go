package objects

type ProductFeedDelete struct {
	Success *bool `json:"success,omitempty"`
}

var ProductFeedDeleteFields = struct {
	Success string
}{
	Success: "success",
}
