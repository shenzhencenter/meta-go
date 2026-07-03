package enums

type OpenbridgeconfigurationEventEnrichmentState string

const (
	OpenbridgeconfigurationEventEnrichmentStateNo             OpenbridgeconfigurationEventEnrichmentState = "NO"
	OpenbridgeconfigurationEventEnrichmentStateNotInitialized OpenbridgeconfigurationEventEnrichmentState = "NOT_INITIALIZED"
	OpenbridgeconfigurationEventEnrichmentStateYes            OpenbridgeconfigurationEventEnrichmentState = "YES"
)

func (value OpenbridgeconfigurationEventEnrichmentState) String() string {
	return string(value)
}
