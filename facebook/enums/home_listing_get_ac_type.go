package enums

type HomelistinggetAcType string

const (
	HomelistinggetAcTypeCentral    HomelistinggetAcType = "CENTRAL"
	HomelistinggetAcTypeEmptyValue HomelistinggetAcType = "EMPTY_VALUE"
	HomelistinggetAcTypeNone       HomelistinggetAcType = "NONE"
	HomelistinggetAcTypeOther      HomelistinggetAcType = "OTHER"
)

func (value HomelistinggetAcType) String() string {
	return string(value)
}
