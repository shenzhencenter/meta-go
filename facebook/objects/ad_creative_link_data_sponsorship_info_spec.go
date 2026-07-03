package objects

type AdCreativeLinkDataSponsorshipInfoSpec struct {
	SponsorImageURL *string `json:"sponsor_image_url,omitempty"`
	SponsorName     *string `json:"sponsor_name,omitempty"`
}

var AdCreativeLinkDataSponsorshipInfoSpecFields = struct {
	SponsorImageURL string
	SponsorName     string
}{
	SponsorImageURL: "sponsor_image_url",
	SponsorName:     "sponsor_name",
}
