package enums

type LivevideoPersistentStreamKeyStatus string

const (
	LivevideoPersistentStreamKeyStatusDisable    LivevideoPersistentStreamKeyStatus = "DISABLE"
	LivevideoPersistentStreamKeyStatusEnable     LivevideoPersistentStreamKeyStatus = "ENABLE"
	LivevideoPersistentStreamKeyStatusRegenerate LivevideoPersistentStreamKeyStatus = "REGENERATE"
)

func (value LivevideoPersistentStreamKeyStatus) String() string {
	return string(value)
}
