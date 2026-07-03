package objects

type AdCreativeStaticFallbackSpec struct {
	CallToAction *AdCreativeLinkDataCallToAction `json:"call_to_action,omitempty"`
	Description  *string                         `json:"description,omitempty"`
	ImageHash    *string                         `json:"image_hash,omitempty"`
	Link         *string                         `json:"link,omitempty"`
	Message      *string                         `json:"message,omitempty"`
	Name         *string                         `json:"name,omitempty"`
}

var AdCreativeStaticFallbackSpecFields = struct {
	CallToAction string
	Description  string
	ImageHash    string
	Link         string
	Message      string
	Name         string
}{
	CallToAction: "call_to_action",
	Description:  "description",
	ImageHash:    "image_hash",
	Link:         "link",
	Message:      "message",
	Name:         "name",
}
