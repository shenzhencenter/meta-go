package enums

type ExtendedcreditowningCreditAllocationConfigsPartitionTypeEnumParam string

const (
	ExtendedcreditowningCreditAllocationConfigsPartitionTypeEnumParamAuth                  ExtendedcreditowningCreditAllocationConfigsPartitionTypeEnumParam = "AUTH"
	ExtendedcreditowningCreditAllocationConfigsPartitionTypeEnumParamFixed                 ExtendedcreditowningCreditAllocationConfigsPartitionTypeEnumParam = "FIXED"
	ExtendedcreditowningCreditAllocationConfigsPartitionTypeEnumParamFixedWithoutPartition ExtendedcreditowningCreditAllocationConfigsPartitionTypeEnumParam = "FIXED_WITHOUT_PARTITION"
)

func (value ExtendedcreditowningCreditAllocationConfigsPartitionTypeEnumParam) String() string {
	return string(value)
}
