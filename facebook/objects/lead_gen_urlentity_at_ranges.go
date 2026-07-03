package objects

type LeadGenURLEntityAtRanges struct {
	Length *uint64 `json:"length,omitempty"`
	Offset *uint64 `json:"offset,omitempty"`
	URL    *string `json:"url,omitempty"`
}

var LeadGenURLEntityAtRangesFields = struct {
	Length string
	Offset string
	URL    string
}{
	Length: "length",
	Offset: "offset",
	URL:    "url",
}
