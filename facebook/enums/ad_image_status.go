package enums

type AdimageStatus string

const (
	AdimageStatusActive   AdimageStatus = "ACTIVE"
	AdimageStatusDeleted  AdimageStatus = "DELETED"
	AdimageStatusInternal AdimageStatus = "INTERNAL"
)

func (value AdimageStatus) String() string {
	return string(value)
}
