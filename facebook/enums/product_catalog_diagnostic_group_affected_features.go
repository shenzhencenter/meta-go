package enums

type ProductcatalogdiagnosticgroupAffectedFeatures string

const (
	ProductcatalogdiagnosticgroupAffectedFeaturesAugmentedReality ProductcatalogdiagnosticgroupAffectedFeatures = "augmented_reality"
	ProductcatalogdiagnosticgroupAffectedFeaturesCheckout         ProductcatalogdiagnosticgroupAffectedFeatures = "checkout"
)

func (value ProductcatalogdiagnosticgroupAffectedFeatures) String() string {
	return string(value)
}
