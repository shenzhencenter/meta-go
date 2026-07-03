package enums

type ProductgroupDeletionMethod string

const (
	ProductgroupDeletionMethodDeleteItems ProductgroupDeletionMethod = "DELETE_ITEMS"
	ProductgroupDeletionMethodOnlyIfEmpty ProductgroupDeletionMethod = "ONLY_IF_EMPTY"
)

func (value ProductgroupDeletionMethod) String() string {
	return string(value)
}
