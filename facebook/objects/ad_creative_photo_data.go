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
