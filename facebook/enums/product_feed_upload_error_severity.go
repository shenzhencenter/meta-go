package enums

type ProductfeeduploaderrorSeverity string

const (
	ProductfeeduploaderrorSeverityFatal   ProductfeeduploaderrorSeverity = "fatal"
	ProductfeeduploaderrorSeverityWarning ProductfeeduploaderrorSeverity = "warning"
)

func (value ProductfeeduploaderrorSeverity) String() string {
	return string(value)
}
