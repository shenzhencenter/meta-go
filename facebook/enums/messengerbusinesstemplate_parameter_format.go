package enums

type MessengerbusinesstemplateParameterFormat string

const (
	MessengerbusinesstemplateParameterFormatNamed      MessengerbusinesstemplateParameterFormat = "NAMED"
	MessengerbusinesstemplateParameterFormatPositional MessengerbusinesstemplateParameterFormat = "POSITIONAL"
)

func (value MessengerbusinesstemplateParameterFormat) String() string {
	return string(value)
}
