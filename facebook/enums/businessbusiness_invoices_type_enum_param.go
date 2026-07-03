package enums

type BusinessbusinessInvoicesTypeEnumParam string

const (
	BusinessbusinessInvoicesTypeEnumParamCm       BusinessbusinessInvoicesTypeEnumParam = "CM"
	BusinessbusinessInvoicesTypeEnumParamDm       BusinessbusinessInvoicesTypeEnumParam = "DM"
	BusinessbusinessInvoicesTypeEnumParamInv      BusinessbusinessInvoicesTypeEnumParam = "INV"
	BusinessbusinessInvoicesTypeEnumParamProForma BusinessbusinessInvoicesTypeEnumParam = "PRO_FORMA"
)

func (value BusinessbusinessInvoicesTypeEnumParam) String() string {
	return string(value)
}
