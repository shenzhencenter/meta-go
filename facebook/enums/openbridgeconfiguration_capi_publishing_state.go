package enums

type OpenbridgeconfigurationCapiPublishingState string

const (
	OpenbridgeconfigurationCapiPublishingStateDisabled       OpenbridgeconfigurationCapiPublishingState = "DISABLED"
	OpenbridgeconfigurationCapiPublishingStateEnabled        OpenbridgeconfigurationCapiPublishingState = "ENABLED"
	OpenbridgeconfigurationCapiPublishingStateNotInitialized OpenbridgeconfigurationCapiPublishingState = "NOT_INITIALIZED"
)

func (value OpenbridgeconfigurationCapiPublishingState) String() string {
	return string(value)
}
