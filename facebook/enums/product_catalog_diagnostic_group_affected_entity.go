package enums

type ProductcatalogdiagnosticgroupAffectedEntity string

const (
	ProductcatalogdiagnosticgroupAffectedEntityProductCatalog ProductcatalogdiagnosticgroupAffectedEntity = "product_catalog"
	ProductcatalogdiagnosticgroupAffectedEntityProductEvent   ProductcatalogdiagnosticgroupAffectedEntity = "product_event"
	ProductcatalogdiagnosticgroupAffectedEntityProductItem    ProductcatalogdiagnosticgroupAffectedEntity = "product_item"
	ProductcatalogdiagnosticgroupAffectedEntityProductSet     ProductcatalogdiagnosticgroupAffectedEntity = "product_set"
)

func (value ProductcatalogdiagnosticgroupAffectedEntity) String() string {
	return string(value)
}
