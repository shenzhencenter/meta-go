package enums

type PrivateliftstudyinstanceOperation string

const (
	PrivateliftstudyinstanceOperationAggregate PrivateliftstudyinstanceOperation = "AGGREGATE"
	PrivateliftstudyinstanceOperationCancel    PrivateliftstudyinstanceOperation = "CANCEL"
	PrivateliftstudyinstanceOperationCompute   PrivateliftstudyinstanceOperation = "COMPUTE"
	PrivateliftstudyinstanceOperationIDMatch   PrivateliftstudyinstanceOperation = "ID_MATCH"
	PrivateliftstudyinstanceOperationNext      PrivateliftstudyinstanceOperation = "NEXT"
	PrivateliftstudyinstanceOperationNone      PrivateliftstudyinstanceOperation = "NONE"
)

func (value PrivateliftstudyinstanceOperation) String() string {
	return string(value)
}
