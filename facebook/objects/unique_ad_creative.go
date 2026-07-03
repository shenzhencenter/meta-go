package objects

type UniqueAdCreative struct {
	SampleCreative *AdCreative `json:"sample_creative,omitempty"`
	VisualHash     *uint64     `json:"visual_hash,omitempty"`
}
