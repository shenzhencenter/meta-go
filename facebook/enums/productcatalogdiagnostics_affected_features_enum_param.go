package enums

type ProductcatalogdiagnosticsAffectedFeaturesEnumParam string

const (
	ProductcatalogdiagnosticsAffectedFeaturesEnumParamAugmentedReality ProductcatalogdiagnosticsAffectedFeaturesEnumParam = "augmented_reality"
	ProductcatalogdiagnosticsAffectedFeaturesEnumParamCheckout         ProductcatalogdiagnosticsAffectedFeaturesEnumParam = "checkout"
)

func (value ProductcatalogdiagnosticsAffectedFeaturesEnumParam) String() string {
	return string(value)
}
