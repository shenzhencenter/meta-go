package enums

type BusinessmanagedPartnerBusinessesPartitionTypeEnumParam string

const (
	BusinessmanagedPartnerBusinessesPartitionTypeEnumParamAuth                  BusinessmanagedPartnerBusinessesPartitionTypeEnumParam = "AUTH"
	BusinessmanagedPartnerBusinessesPartitionTypeEnumParamFixed                 BusinessmanagedPartnerBusinessesPartitionTypeEnumParam = "FIXED"
	BusinessmanagedPartnerBusinessesPartitionTypeEnumParamFixedWithoutPartition BusinessmanagedPartnerBusinessesPartitionTypeEnumParam = "FIXED_WITHOUT_PARTITION"
)

func (value BusinessmanagedPartnerBusinessesPartitionTypeEnumParam) String() string {
	return string(value)
}
