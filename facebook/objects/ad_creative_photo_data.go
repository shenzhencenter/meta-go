package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativePhotoData struct {
	BrandedContentSharedToSponsorStatus *string  `json:"branded_content_shared_to_sponsor_status,omitempty"`
	BrandedContentSponsorPageID         *core.ID `json:"branded_content_sponsor_page_id,omitempty"`
	Caption                             *string  `json:"caption,omitempty"`
	ImageHash                           *string  `json:"image_hash,omitempty"`
	PageWelcomeMessage                  *string  `json:"page_welcome_message,omitempty"`
	URL                                 *string  `json:"url,omitempty"`
}

var AdCreativePhotoDataFields = struct {
	BrandedContentSharedToSponsorStatus string
	BrandedContentSponsorPageID         string
	Caption                             string
	ImageHash                           string
	PageWelcomeMessage                  string
	URL                                 string
}{
	BrandedContentSharedToSponsorStatus: "branded_content_shared_to_sponsor_status",
	BrandedContentSponsorPageID:         "branded_content_sponsor_page_id",
	Caption:                             "caption",
	ImageHash:                           "image_hash",
	PageWelcomeMessage:                  "page_welcome_message",
	URL:                                 "url",
}
