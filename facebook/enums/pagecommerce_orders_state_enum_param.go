package enums

type PagecommerceOrdersStateEnumParam string

const (
	PagecommerceOrdersStateEnumParamCompleted    PagecommerceOrdersStateEnumParam = "COMPLETED"
	PagecommerceOrdersStateEnumParamCreated      PagecommerceOrdersStateEnumParam = "CREATED"
	PagecommerceOrdersStateEnumParamFbProcessing PagecommerceOrdersStateEnumParam = "FB_PROCESSING"
	PagecommerceOrdersStateEnumParamInProgress   PagecommerceOrdersStateEnumParam = "IN_PROGRESS"
)

func (value PagecommerceOrdersStateEnumParam) String() string {
	return string(value)
}
