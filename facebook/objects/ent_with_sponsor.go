package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type EntWithSponsor struct {
	HasAudioSwappedFbCopy        *bool    `json:"has_audio_swapped_fb_copy,omitempty"`
	ID                           *core.ID `json:"id,omitempty"`
	IsBrandedContent             *bool    `json:"is_branded_content,omitempty"`
	OwnerLinkedInstagramUserV1ID *core.ID `json:"owner_linked_instagram_user_v1_id,omitempty"`
	OwnerPicture                 *string  `json:"owner_picture,omitempty"`
	PostID                       *core.ID `json:"post_id,omitempty"`
	PostInfo                     *Post    `json:"post_info,omitempty"`
	PromotableID                 *core.ID `json:"promotable_id,omitempty"`
}

var EntWithSponsorFields = struct {
	HasAudioSwappedFbCopy        string
	ID                           string
	IsBrandedContent             string
	OwnerLinkedInstagramUserV1ID string
	OwnerPicture                 string
	PostID                       string
	PostInfo                     string
	PromotableID                 string
}{
	HasAudioSwappedFbCopy:        "has_audio_swapped_fb_copy",
	ID:                           "id",
	IsBrandedContent:             "is_branded_content",
	OwnerLinkedInstagramUserV1ID: "owner_linked_instagram_user_v1_id",
	OwnerPicture:                 "owner_picture",
	PostID:                       "post_id",
	PostInfo:                     "post_info",
	PromotableID:                 "promotable_id",
}
