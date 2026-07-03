package enums

type OpenbridgeconfigurationEventEnrichmentAdvertiserState string

const (
	OpenbridgeconfigurationEventEnrichmentAdvertiserStateDisabled       OpenbridgeconfigurationEventEnrichmentAdvertiserState = "DISABLED"
	OpenbridgeconfigurationEventEnrichmentAdvertiserStateEnabled        OpenbridgeconfigurationEventEnrichmentAdvertiserState = "ENABLED"
	OpenbridgeconfigurationEventEnrichmentAdvertiserStateNotInitialized OpenbridgeconfigurationEventEnrichmentAdvertiserState = "NOT_INITIALIZED"
)

func (value OpenbridgeconfigurationEventEnrichmentAdvertiserState) String() string {
	return string(value)
}
