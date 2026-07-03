package enums

type GroupIconSize string

const (
	GroupIconSizeX16 GroupIconSize = "16"
	GroupIconSizeX34 GroupIconSize = "34"
	GroupIconSizeX50 GroupIconSize = "50"
	GroupIconSizeX68 GroupIconSize = "68"
)

func (value GroupIconSize) String() string {
	return string(value)
}
