package objects

type AdCreativeStaticFallbackSpec struct {
	CallToAction *AdCreativeLinkDataCallToAction `json:"call_to_action,omitempty"`
	Description  *string                         `json:"description,omitempty"`
	ImageHash    *string                         `json:"image_hash,omitempty"`
	Link         *string                         `json:"link,omitempty"`
	Message      *string                         `json:"message,omitempty"`
	Name         *string                         `json:"name,omitempty"`
}
