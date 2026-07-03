package enums

type LivevideoStreamType string

const (
	LivevideoStreamTypeAmbient LivevideoStreamType = "AMBIENT"
	LivevideoStreamTypeRegular LivevideoStreamType = "REGULAR"
)

func (value LivevideoStreamType) String() string {
	return string(value)
}
