package enums

type OpenbridgeconfigurationEventEnrichmentMetaState string

const (
	OpenbridgeconfigurationEventEnrichmentMetaStateAllowed        OpenbridgeconfigurationEventEnrichmentMetaState = "ALLOWED"
	OpenbridgeconfigurationEventEnrichmentMetaStateBlocked        OpenbridgeconfigurationEventEnrichmentMetaState = "BLOCKED"
	OpenbridgeconfigurationEventEnrichmentMetaStateNotInitialized OpenbridgeconfigurationEventEnrichmentMetaState = "NOT_INITIALIZED"
)

func (value OpenbridgeconfigurationEventEnrichmentMetaState) String() string {
	return string(value)
}
