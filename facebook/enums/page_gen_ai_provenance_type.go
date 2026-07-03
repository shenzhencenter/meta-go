package enums

type PageGenAiProvenanceType string

const (
	PageGenAiProvenanceTypeC2pa               PageGenAiProvenanceType = "C2PA"
	PageGenAiProvenanceTypeC2paMetadataEdited PageGenAiProvenanceType = "C2PA_METADATA_EDITED"
	PageGenAiProvenanceTypeExplicit           PageGenAiProvenanceType = "EXPLICIT"
	PageGenAiProvenanceTypeExplicitAnimate    PageGenAiProvenanceType = "EXPLICIT_ANIMATE"
	PageGenAiProvenanceTypeExplicitDropIn     PageGenAiProvenanceType = "EXPLICIT_DROP_IN"
	PageGenAiProvenanceTypeExplicitFaceSwap   PageGenAiProvenanceType = "EXPLICIT_FACE_SWAP"
	PageGenAiProvenanceTypeExplicitImagine    PageGenAiProvenanceType = "EXPLICIT_IMAGINE"
	PageGenAiProvenanceTypeExplicitImagineMe  PageGenAiProvenanceType = "EXPLICIT_IMAGINE_ME"
	PageGenAiProvenanceTypeExplicitRestyle    PageGenAiProvenanceType = "EXPLICIT_RESTYLE"
	PageGenAiProvenanceTypeExplicitWardrobe   PageGenAiProvenanceType = "EXPLICIT_WARDROBE"
	PageGenAiProvenanceTypeInvisibleWatermark PageGenAiProvenanceType = "INVISIBLE_WATERMARK"
	PageGenAiProvenanceTypeIptc               PageGenAiProvenanceType = "IPTC"
	PageGenAiProvenanceTypeIptcMetadataEdited PageGenAiProvenanceType = "IPTC_METADATA_EDITED"
)

func (value PageGenAiProvenanceType) String() string {
	return string(value)
}
